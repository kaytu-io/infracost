package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getDXConnectionRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_dx_connection",
		RFunc: NewDXConnection,
	}
}

func NewDXConnection(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.DXConnection{
		Address:   d.Address,
		Region:    d.Get("region").String(),
		Bandwidth: d.Get("bandwidth").String(),
		Location:  d.Get("location").String(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
