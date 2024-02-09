package terraform

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/kaytu-io/infracost/external/clierror"
	"github.com/kaytu-io/infracost/external/config"
	"github.com/kaytu-io/infracost/external/schema"
)

type PlanProvider struct {
	*DirProvider
	Path                 string
	cachedPlanJSON       []byte
	includePastResources bool
}

func NewPlanProvider(ctx *config.ProjectContext, includePastResources bool) schema.Provider {
	dirProvider := NewDirProvider(ctx, includePastResources).(*DirProvider)

	return &PlanProvider{
		DirProvider:          dirProvider,
		Path:                 ctx.ProjectConfig.Path,
		includePastResources: includePastResources,
	}
}

func (p *PlanProvider) Type() string {
	return "terraform_plan_binary"
}

func (p *PlanProvider) DisplayType() string {
	return "Terraform plan binary file"
}

func (p *PlanProvider) LoadResources(usage schema.UsageMap) ([]*schema.Project, error) {
	j, err := p.generatePlanJSON()
	if err != nil {
		return []*schema.Project{}, err
	}

	metadata := config.DetectProjectMetadata(p.ctx.ProjectConfig.Path)
	metadata.Type = p.Type()
	p.AddMetadata(metadata)
	name := p.ctx.ProjectConfig.Name

	project := schema.NewProject(name, metadata)
	parser := NewParser(p.ctx, p.includePastResources)

	partialPastResources, partialResources, providerMetadatas, err := parser.parseJSON(j, usage)
	if err != nil {
		return []*schema.Project{project}, errors.Wrap(err, "Error parsing Terraform JSON")
	}

	project.AddProviderMetadata(providerMetadatas)

	project.PartialPastResources = partialPastResources
	project.PartialResources = partialResources

	return []*schema.Project{project}, nil
}

func (p *PlanProvider) generatePlanJSON() ([]byte, error) {
	if p.cachedPlanJSON != nil {
		return p.cachedPlanJSON, nil
	}

	dir := filepath.Dir(p.Path)
	planPath := filepath.Base(p.Path)

	if !IsTerraformDir(dir) {
		log.Debug().Msgf("%s is not a Terraform directory, checking current working directory", dir)
		dir, err := os.Getwd()
		if err != nil {
			return []byte{}, err
		}
		planPath = p.Path

		if !IsTerraformDir(dir) {
			m := fmt.Sprintf("%s %s.",
				"Could not detect Terraform directory for",
				p.Path,
			)
			return []byte{}, clierror.NewCLIError(errors.New(m), "Could not detect Terraform directory for plan file")
		}
	}

	err := p.checks()
	if err != nil {
		return []byte{}, err
	}

	opts, err := p.buildCommandOpts(dir)
	if err != nil {
		return []byte{}, err
	}
	if opts.TerraformConfigFile != "" {
		defer os.Remove(opts.TerraformConfigFile)
	}

	j, err := p.runShow(opts, planPath, false)
	if err == nil {
		p.cachedPlanJSON = j
	}
	return j, err
}
