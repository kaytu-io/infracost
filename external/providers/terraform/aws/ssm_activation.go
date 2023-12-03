package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getSSMActivationRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_ssm_activation",
		RFunc: NewSSMActivation,
	}
}

func NewSSMActivation(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.SSMActivation{
		Address:           d.Address,
		Region:            d.Get("region").String(),
		RegistrationLimit: d.Get("registration_limit").Int(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
