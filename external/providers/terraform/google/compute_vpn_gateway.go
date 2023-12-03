package google

import (
	"github.com/kaytu-io/infracost/external/resources/google"
	"github.com/kaytu-io/infracost/external/schema"
)

func getComputeVPNGatewayRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "google_compute_vpn_gateway",
		RFunc: NewComputeVPNGateway,
	}
}
func NewComputeVPNGateway(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &google.ComputeVPNGateway{Address: d.Address, Region: d.Get("region").String()}
	r.PopulateUsage(u)
	return r.BuildResource()
}
