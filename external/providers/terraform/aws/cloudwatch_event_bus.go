package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getCloudwatchEventBusItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_cloudwatch_event_bus",
		RFunc: NewCloudwatchEventBus,
	}
}
func NewCloudwatchEventBus(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.CloudwatchEventBus{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
