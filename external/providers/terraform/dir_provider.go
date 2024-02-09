package terraform

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/kballard/go-shellquote"
	"github.com/pkg/errors"
	"golang.org/x/mod/semver"

	"github.com/kaytu-io/infracost/external/clierror"
	"github.com/kaytu-io/infracost/external/config"
	"github.com/kaytu-io/infracost/external/credentials"
	"github.com/kaytu-io/infracost/external/schema"
	"github.com/rs/zerolog/log"
)

var minTerraformVer = "v0.12"

type DirProvider struct {
	ctx                  *config.ProjectContext
	Path                 string
	IsTerragrunt         bool
	PlanFlags            string
	InitFlags            string
	Workspace            string
	UseState             bool
	TerraformBinary      string
	TerraformCloudHost   string
	TerraformCloudToken  string
	Env                  map[string]string
	cachedStateJSON      []byte
	cachedPlanJSON       []byte
	includePastResources bool
}

type RunShowOptions struct {
	CmdOptions *CmdOptions
}

func NewDirProvider(ctx *config.ProjectContext, includePastResources bool) schema.Provider {
	terraformBinary := ctx.ProjectConfig.TerraformBinary
	if terraformBinary == "" {
		terraformBinary = defaultTerraformBinary
	}

	return &DirProvider{
		ctx:                  ctx,
		Path:                 ctx.ProjectConfig.Path,
		PlanFlags:            ctx.ProjectConfig.TerraformPlanFlags,
		InitFlags:            ctx.ProjectConfig.TerraformInitFlags,
		Workspace:            ctx.ProjectConfig.TerraformWorkspace,
		UseState:             ctx.ProjectConfig.TerraformUseState,
		TerraformBinary:      terraformBinary,
		TerraformCloudHost:   ctx.ProjectConfig.TerraformCloudHost,
		TerraformCloudToken:  ctx.ProjectConfig.TerraformCloudToken,
		Env:                  ctx.ProjectConfig.Env,
		includePastResources: includePastResources,
	}
}

func (p *DirProvider) Type() string {
	return "terraform_cli"
}

func (p *DirProvider) DisplayType() string {
	return "Terraform CLI"
}

func (p *DirProvider) checks() error {
	binary := p.TerraformBinary

	p.ctx.ContextValues.SetValue("terraformBinary", binary)

	_, err := exec.LookPath(binary)
	if err != nil {
		msg := fmt.Sprintf("Terraform binary '%s' could not be found. You have two options:\n", binary)
		msg += "1. Set a custom Terraform binary using the environment variable INFRACOST_TERRAFORM_BINARY.\n\n"
		return clierror.NewCLIError(errors.Errorf(msg), "Terraform binary could not be found")
	}

	out, err := exec.Command(binary, "-version").Output()
	if err != nil {
		msg := fmt.Sprintf("Could not get version of Terraform binary '%s'", binary)
		return clierror.NewCLIError(errors.Errorf(msg), "Could not get version of Terraform binary")
	}

	fullVersion := strings.SplitN(string(out), "\n", 2)[0]
	version := shortTerraformVersion(fullVersion)

	p.ctx.ContextValues.SetValue("terraformFullVersion", fullVersion)
	p.ctx.ContextValues.SetValue("terraformVersion", version)

	return checkTerraformVersion(version, fullVersion)
}

func (p *DirProvider) AddMetadata(metadata *schema.ProjectMetadata) {
	metadata.ConfigSha = p.ctx.ProjectConfig.ConfigSha

	basePath := p.ctx.ProjectConfig.Path
	if p.ctx.RunContext.Config.ConfigFilePath != "" {
		basePath = filepath.Dir(p.ctx.RunContext.Config.ConfigFilePath)
	}

	modulePath, err := filepath.Rel(basePath, metadata.Path)
	if err == nil && modulePath != "" && modulePath != "." {
		log.Debug().Msgf("Calculated relative terraformModulePath for %s from %s", basePath, metadata.Path)
		metadata.TerraformModulePath = modulePath
	}

	terraformWorkspace := p.Workspace

	if terraformWorkspace == "" {
		binary := p.TerraformBinary
		cmd := exec.Command(binary, "workspace", "show")
		cmd.Dir = p.Path

		out, err := cmd.Output()
		if err != nil {
			log.Debug().Msgf("Could not detect Terraform workspace for %s", p.Path)
		}
		terraformWorkspace = strings.Split(string(out), "\n")[0]
	}

	metadata.TerraformWorkspace = terraformWorkspace
}

func (p *DirProvider) LoadResources(usage schema.UsageMap) ([]*schema.Project, error) {
	projects := make([]*schema.Project, 0)
	var out []byte
	var err error

	if p.UseState {
		out, err = p.generateStateJSON()
	} else {
		out, err = p.generatePlanJSON()
	}
	if err != nil {
		return projects, err
	}

	jsons := [][]byte{out}
	if p.IsTerragrunt {
		jsons = bytes.Split(out, []byte{'\n'})
		if len(jsons) > 1 {
			jsons = jsons[:len(jsons)-1]
		}
	}

	for _, j := range jsons {
		metadata := config.DetectProjectMetadata(p.ctx.ProjectConfig.Path)
		metadata.Type = p.Type()
		p.AddMetadata(metadata)
		name := p.ctx.ProjectConfig.Name

		project := schema.NewProject(name, metadata)

		parser := NewParser(p.ctx, p.includePastResources)
		pastpartialResources, partialResources, providerMetadatas, err := parser.parseJSON(j, usage)
		if err != nil {
			return projects, errors.Wrap(err, "Error parsing Terraform JSON")
		}

		project.AddProviderMetadata(providerMetadatas)

		project.HasDiff = !p.UseState
		if project.HasDiff {
			project.PartialPastResources = pastpartialResources
		}
		project.PartialResources = partialResources

		projects = append(projects, project)
	}

	return projects, nil
}

func (p *DirProvider) generatePlanJSON() ([]byte, error) {
	if p.cachedPlanJSON != nil {
		return p.cachedPlanJSON, nil
	}

	if UsePlanCache(p) {

		cached, err := ReadPlanCache(p)
		if err != nil {
		} else {
			p.cachedPlanJSON = cached
			return p.cachedPlanJSON, nil
		}
	}

	err := p.checks()
	if err != nil {
		return []byte{}, err
	}

	opts, err := p.buildCommandOpts(p.Path)
	if err != nil {
		return []byte{}, err
	}
	if opts.TerraformConfigFile != "" {
		defer os.Remove(opts.TerraformConfigFile)
	}

	planFile, planJSON, err := p.runPlan(opts, true)
	defer os.Remove(planFile)

	if err != nil {
		return []byte{}, err
	}

	if len(planJSON) > 0 {
		return planJSON, nil
	}

	j, err := p.runShow(opts, planFile, false)
	if err == nil {
		p.cachedPlanJSON = j
		if UsePlanCache(p) {
			// Note we check UsePlanCache again because we have discovered we're using remote execution inside p.runPlan
			WritePlanCache(p, j)
		}
	}
	return j, err
}

func (p *DirProvider) generateStateJSON() ([]byte, error) {
	if p.cachedStateJSON != nil {
		return p.cachedStateJSON, nil
	}

	err := p.checks()
	if err != nil {
		return []byte{}, err
	}

	opts, err := p.buildCommandOpts(p.Path)
	if err != nil {
		return []byte{}, err
	}
	if opts.TerraformConfigFile != "" {
		defer os.Remove(opts.TerraformConfigFile)
	}

	j, err := p.runShow(opts, "", true)
	if err == nil {
		p.cachedStateJSON = j
	}
	return j, err
}

func (p *DirProvider) buildCommandOpts(path string) (*CmdOptions, error) {
	opts := &CmdOptions{
		TerraformBinary:    p.TerraformBinary,
		TerraformWorkspace: p.Workspace,
		Dir:                path,
		Env:                p.Env,
	}

	cfgFile, err := CreateConfigFile(p.Path, p.TerraformCloudHost, p.TerraformCloudToken)
	if err != nil {
		return opts, err
	}

	opts.TerraformConfigFile = cfgFile

	return opts, nil
}

func (p *DirProvider) runPlan(opts *CmdOptions, initOnFail bool) (string, []byte, error) {
	var planJSON []byte

	fileName := ".tfplan-" + uuid.New().String()
	// For Terragrunt we need a relative path
	if !p.IsTerragrunt {
		fileName = filepath.Join(os.TempDir(), fileName)
	}

	flags, err := shellquote.Split(p.PlanFlags)
	if err != nil {
		return "", planJSON, errors.Wrap(err, "Error parsing terraform plan flags")
	}

	args := []string{}
	if p.IsTerragrunt {
		args = append(args, "run-all", "--terragrunt-ignore-external-dependencies")
	}

	args = append(args, "plan", "-input=false", "-lock=false", "-no-color")
	args = append(args, flags...)
	_, err = Cmd(opts, append(args, fmt.Sprintf("-out=%s", fileName))...)

	// Check if the error requires a remote run or an init
	if err != nil {
		extractedErr := extractStderr(err)

		// If the plan returns this error then Terraform is configured with remote execution mode
		if isTerraformRemoteExecutionErr(extractedErr) {
			log.Info().Msg("Continuing with Terraform Remote Execution Mode")
			p.ctx.ContextValues.SetValue("terraformRemoteExecutionModeEnabled", true)
			planJSON, err = p.runRemotePlan(opts, args)
		} else if initOnFail && isTerraformInitErr(extractedErr) {
			err = p.runInit(opts)
			if err != nil {
				return "", planJSON, err
			}
			return p.runPlan(opts, false)
		}
	}

	if err != nil {
		cmdName := "terraform plan"
		if p.IsTerragrunt {
			cmdName = "terragrunt run-all plan"
		}
		msg := fmt.Sprintf("%s failed", cmdName)
		return "", planJSON, clierror.NewCLIError(fmt.Errorf("%s: %s", msg, err), msg)
	}

	return fileName, planJSON, nil
}

func (p *DirProvider) runInit(opts *CmdOptions) error {
	args := []string{}
	if p.IsTerragrunt {
		args = append(args, "run-all", "--terragrunt-ignore-external-dependencies")
	}

	flags, err := shellquote.Split(p.InitFlags)
	if err != nil {
		msg := "parsing terraform-init-flags failed"
		return clierror.NewCLIError(fmt.Errorf("%s: %s", msg, err), msg)
	}

	args = append(args, "init", "-input=false", "-no-color")
	args = append(args, flags...)

	if config.IsTest() {
		args = append(args, "-upgrade")
	}

	_, err = Cmd(opts, args...)
	if err != nil {
		cmdName := "terraform init"
		if p.IsTerragrunt {
			cmdName = "terragrunt run-all init"
		}
		msg := fmt.Sprintf("%s failed", cmdName)
		return clierror.NewCLIError(fmt.Errorf("%s: %s", msg, err), msg)
	}

	return nil
}

func (p *DirProvider) runRemotePlan(opts *CmdOptions, args []string) ([]byte, error) {
	if p.TerraformCloudToken == "" && !credentials.CheckCloudConfigSet() {
		return []byte{}, credentials.ErrMissingCloudToken
	}

	stdout, err := Cmd(opts, args...)
	if err != nil {
		return []byte{}, err
	}

	r := regexp.MustCompile(`To view this run in a browser, visit:\n(.*)`)
	matches := r.FindAllStringSubmatch(string(stdout), 1)
	if len(matches) == 0 || len(matches[0]) <= 1 {
		return []byte{}, errors.New("Could not parse the remote run URL")
	}

	u, err := url.Parse(matches[0][1])
	if err != nil {
		return []byte{}, err
	}
	host := u.Host
	s := strings.Split(u.Path, "/")
	runID := s[len(s)-1]

	token := p.TerraformCloudToken
	if token == "" {
		token = credentials.FindTerraformCloudToken(host)
	}
	if token == "" {
		return []byte{}, credentials.ErrMissingCloudToken
	}

	body, err := cloudAPI(host, fmt.Sprintf("/api/v2/runs/%s/plan", runID), token)
	if err != nil {
		return []byte{}, err
	}

	var parsedResp struct {
		Data struct {
			Links map[string]string
		}
	}
	if json.Unmarshal(body, &parsedResp) != nil {
		return []byte{}, err
	}

	jsonPath, ok := parsedResp.Data.Links["json-output"]
	if !ok || jsonPath == "" {
		return []byte{}, errors.New("Could not parse path to plan JSON from remote")
	}
	return cloudAPI(host, jsonPath, token)
}

func (p *DirProvider) runShow(opts *CmdOptions, planFile string, initOnFail bool) ([]byte, error) {
	args := []string{"show", "-no-color", "-json"}
	if planFile != "" {
		args = append(args, planFile)
	}
	out, err := Cmd(opts, args...)

	// Check if the error requires a remote run or an init
	if err != nil {
		extractedErr := extractStderr(err)

		// If the plan returns this error then Terraform is configured with remote execution mode
		if isTerraformRemoteExecutionErr(extractedErr) {
			log.Info().Msg("Terraform expected Remote Execution Mode")
		} else if initOnFail && isTerraformInitErr(extractedErr) {
			err = p.runInit(opts)
			if err != nil {
				return out, err
			}
			return p.runShow(opts, planFile, false)
		}
	}

	if err != nil {
		cmdName := "terraform show"
		if p.IsTerragrunt {
			cmdName = "terragrunt show"
		}
		msg := fmt.Sprintf("%s failed", cmdName)
		return []byte{}, clierror.NewCLIError(fmt.Errorf("%s: %s", msg, err), msg)
	}

	return out, nil
}

func IsTerraformDir(path string) bool {
	for _, ext := range []string{"tf", "tf.json"} {
		matches, err := filepath.Glob(filepath.Join(path, fmt.Sprintf("*.%s", ext)))
		if matches != nil && err == nil {
			return true
		}
	}
	return false
}

func shortTerraformVersion(full string) string {
	p := strings.Split(full, " ")
	if len(p) > 1 {
		return p[len(p)-1]
	}

	return ""
}

func checkTerraformVersion(v string, fullV string) error {
	if len(v) > 0 && v[0] != 'v' {
		// The semver package requires a 'v' prefix to do a proper Compare.
		v = "v" + v
	}

	if strings.HasPrefix(fullV, "Terraform ") && semver.Compare(v, minTerraformVer) < 0 {
		return fmt.Errorf("Terraform %s is not supported. Please use Terraform version >= %s. Update it or set the environment variable INFRACOST_TERRAFORM_BINARY.", v, minTerraformVer) //nolint
	}

	if strings.HasPrefix(fullV, "terragrunt") && semver.Compare(v, minTerragruntVer) < 0 {
		return fmt.Errorf("Terragrunt %s is not supported. Please use Terragrunt version >= %s. Update it or set the environment variable INFRACOST_TERRAFORM_BINARY.", v, minTerragruntVer) //nolint
	}

	// Allow any non-terraform and non-terragrunt binaries
	return nil
}

func extractStderr(err error) string {
	if e, ok := err.(*CmdError); ok {
		return stripBlankLines(string(e.Stderr))
	}
	return ""
}

func isTerraformRemoteExecutionErr(extractedErr string) bool {
	return strings.HasPrefix(extractedErr, "Error: Saving a generated plan is currently not supported")
}

func isTerraformInitErr(extractedErr string) bool {
	return strings.Contains(extractedErr, "Error: Could not load plugin") ||
		strings.Contains(extractedErr, "Error: Required plugins are not installed") ||
		strings.Contains(extractedErr, "Error: Initialization required") ||
		strings.Contains(extractedErr, "Error: Backend initialization required") ||
		strings.Contains(extractedErr, "Error: Provider requirements cannot be satisfied by locked dependencies") ||
		strings.Contains(extractedErr, "Error: Inconsistent dependency lock file") ||
		strings.Contains(extractedErr, "Error: Module not installed") ||
		strings.Contains(extractedErr, "Error: Terraform Cloud initialization required") ||
		strings.Contains(extractedErr, "please run \"terraform init\"")
}

func stripBlankLines(s string) string {
	return regexp.MustCompile(`[\t\r\n]+`).ReplaceAllString(strings.TrimSpace(s), "\n")
}
