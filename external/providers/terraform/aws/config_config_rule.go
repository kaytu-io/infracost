package aws

import (
	"github.com/infracost/infracost/external/resources/aws"
	"github.com/infracost/infracost/external/schema"
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
