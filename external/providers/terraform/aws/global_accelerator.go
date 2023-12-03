package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getGlobalAcceleratorRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_globalaccelerator_accelerator",
		RFunc: newGlobalAccelerator,
	}
}

func newGlobalAccelerator(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {

	r := &aws.GlobalAccelerator{
		Address: d.Address,
	}

	return r.BuildResource()
}
