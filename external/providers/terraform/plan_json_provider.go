package terraform

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"

	"github.com/kaytu-io/infracost/external/config"
	"github.com/kaytu-io/infracost/external/schema"
)

type PlanJSONProvider struct {
	ctx                  *config.ProjectContext
	Path                 string
	includePastResources bool
	logger               zerolog.Logger
}

func NewPlanJSONProvider(ctx *config.ProjectContext, includePastResources bool) *PlanJSONProvider {
	return &PlanJSONProvider{
		ctx:                  ctx,
		Path:                 ctx.ProjectConfig.Path,
		includePastResources: includePastResources,
		logger:               ctx.Logger(),
	}
}

func (p *PlanJSONProvider) Type() string {
	return "terraform_plan_json"
}

func (p *PlanJSONProvider) DisplayType() string {
	return "Terraform plan JSON file"
}

func (p *PlanJSONProvider) AddMetadata(metadata *schema.ProjectMetadata) {
	metadata.ConfigSha = p.ctx.ProjectConfig.ConfigSha

	// TerraformWorkspace isn't used to load resources but we still pass it
	// on so it appears in the project name of the output
	metadata.TerraformWorkspace = p.ctx.ProjectConfig.TerraformWorkspace
}

func (p *PlanJSONProvider) LoadResources(usage schema.UsageMap) ([]*schema.Project, error) {

	j, err := os.ReadFile(p.Path)
	if err != nil {
		return []*schema.Project{}, fmt.Errorf("Error reading Terraform plan JSON file %w", err)
	}

	project, err := p.LoadResourcesFromSrc(usage, j)
	if err != nil {
		return nil, err
	}

	return []*schema.Project{project}, nil
}

func (p *PlanJSONProvider) LoadResourcesFromSrc(usage schema.UsageMap, j []byte) (*schema.Project, error) {
	metadata := config.DetectProjectMetadata(p.ctx.ProjectConfig.Path)
	metadata.Type = p.Type()
	p.AddMetadata(metadata)
	name := p.ctx.ProjectConfig.Name

	project := schema.NewProject(name, metadata)
	parser := NewParser(p.ctx, p.includePastResources)

	partialPastResources, partialResources, providerMetadatas, err := parser.parseJSON(j, usage)
	if err != nil {
		return project, fmt.Errorf("Error parsing Terraform plan JSON file %w", err)
	}

	project.AddProviderMetadata(providerMetadatas)

	project.PartialPastResources = partialPastResources
	project.PartialResources = partialResources

	return project, nil
}
