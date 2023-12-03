package google

import (
	"github.com/kaytu-io/infracost/external/resources/google"
	"github.com/kaytu-io/infracost/external/schema"
)

func getComputeHAVPNGatewayRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "google_compute_ha_vpn_gateway",
		RFunc: NewComputeHAVPNGateway,
	}
}
func NewComputeHAVPNGateway(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &google.ComputeVPNGateway{Address: d.Address, Region: d.Get("region").String()}
	r.PopulateUsage(u)
	return r.BuildResource()
}
