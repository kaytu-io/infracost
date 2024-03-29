package terraform

import (
	"os"

	"github.com/kaytu-io/infracost/external/config"
	"github.com/kaytu-io/infracost/external/schema"
	"github.com/pkg/errors"
)

type StateJSONProvider struct {
	ctx                  *config.ProjectContext
	Path                 string
	includePastResources bool
}

func NewStateJSONProvider(ctx *config.ProjectContext, includePastResources bool) schema.Provider {
	return &StateJSONProvider{
		ctx:                  ctx,
		Path:                 ctx.ProjectConfig.Path,
		includePastResources: includePastResources,
	}
}

func (p *StateJSONProvider) Type() string {
	return "terraform_state_json"
}

func (p *StateJSONProvider) DisplayType() string {
	return "Terraform state JSON file"
}

func (p *StateJSONProvider) AddMetadata(metadata *schema.ProjectMetadata) {
	metadata.ConfigSha = p.ctx.ProjectConfig.ConfigSha
}

func (p *StateJSONProvider) LoadResources(usage schema.UsageMap) ([]*schema.Project, error) {

	j, err := os.ReadFile(p.Path)
	if err != nil {
		return []*schema.Project{}, errors.Wrap(err, "Error reading Terraform state JSON file")
	}

	metadata := config.DetectProjectMetadata(p.ctx.ProjectConfig.Path)
	metadata.Type = p.Type()
	p.AddMetadata(metadata)
	name := p.ctx.ProjectConfig.Name

	project := schema.NewProject(name, metadata)
	parser := NewParser(p.ctx, p.includePastResources)

	partialPastResources, partialResources, providerMetadatas, err := parser.parseJSON(j, usage)
	if err != nil {
		return []*schema.Project{project}, errors.Wrap(err, "Error parsing Terraform state JSON file")
	}

	project.AddProviderMetadata(providerMetadatas)

	project.PartialPastResources = partialPastResources
	project.PartialResources = partialResources

	return []*schema.Project{project}, nil
}
