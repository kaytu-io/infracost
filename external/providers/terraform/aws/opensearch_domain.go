package aws

import (
	"github.com/infracost/infracost/external/schema"
)

func getOpensearchDomainRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_opensearch_domain",
		RFunc: newSearchDomain,
	}
}
