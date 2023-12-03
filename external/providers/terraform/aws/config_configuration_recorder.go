package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getConfigurationRecorderItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_config_configuration_recorder",
		RFunc: NewConfigConfigurationRecorder,
	}
}
func NewConfigConfigurationRecorder(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.ConfigConfigurationRecorder{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
