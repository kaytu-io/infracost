package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getVPNConnectionRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_vpn_connection",
		RFunc: NewVPNConnection,
	}
}
func NewVPNConnection(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.VPNConnection{Address: d.Address, TransitGatewayID: d.Get("transit_gateway_id").String(), Region: d.Get("region").String()}
	r.PopulateUsage(u)
	return r.BuildResource()
}
