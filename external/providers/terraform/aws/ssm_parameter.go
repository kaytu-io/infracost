package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getSSMParameterRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_ssm_parameter",
		RFunc: NewSSMParameter,
	}
}

func NewSSMParameter(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.SSMParameter{
		Address: d.Address,
		Region:  d.Get("region").String(),
		Tier:    d.Get("tier").String(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
