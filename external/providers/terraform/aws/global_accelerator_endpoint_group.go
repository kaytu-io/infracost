package aws

import (
	"github.com/infracost/infracost/external/resources/aws"
	"github.com/infracost/infracost/external/schema"
)

func getGlobalacceleratorEndpointGroupRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_globalaccelerator_endpoint_group",
		RFunc: newGlobalacceleratorEndpointGroup,
	}
}

func newGlobalacceleratorEndpointGroup(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	region := d.Get("endpoint_group_region").String()
	r := &aws.GlobalacceleratorEndpointGroup{
		Address: d.Address,
		Region:  region,
	}
	r.PopulateUsage(u)

	return r.BuildResource()
}
