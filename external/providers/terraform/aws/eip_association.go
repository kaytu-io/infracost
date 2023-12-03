package aws

import (
	"github.com/infracost/infracost/external/schema"
)

func getEIPAssociationRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:    "aws_eip_association",
		NoPrice: true,
		ReferenceAttributes: []string{
			"allocation_id",
		},
		Notes: []string{"Free resource."},
	}
}
