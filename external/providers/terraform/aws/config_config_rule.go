package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getConfigRuleItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_config_config_rule",
		RFunc: NewConfigConfigRule,
	}
}
func NewConfigConfigRule(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.ConfigConfigRule{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
