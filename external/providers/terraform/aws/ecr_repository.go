package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
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
