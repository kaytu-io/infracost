package aws

import (
	"github.com/infracost/infracost/external/resources/aws"
	"github.com/infracost/infracost/external/schema"
)

func getECRRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:                "aws_ecr_repository",
		CoreRFunc:           NewECRRepository,
		ReferenceAttributes: []string{"aws_ecr_lifecycle_policy.repository"},
	}
}
func NewECRRepository(d *schema.ResourceData) schema.CoreResource {
	return &aws.ECRRepository{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}
}
