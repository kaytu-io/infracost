package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getConfigOrganizationManagedRuleItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_config_organization_managed_rule",
		RFunc: NewConfigOrganizationManagedRule,
	}
}
func NewConfigOrganizationManagedRule(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.ConfigConfigRule{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
