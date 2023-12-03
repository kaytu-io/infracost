package aws

import (
	"github.com/kaytu-io/infracost/external/schema"
)

func getOpensearchDomainRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_opensearch_domain",
		RFunc: newSearchDomain,
	}
}
