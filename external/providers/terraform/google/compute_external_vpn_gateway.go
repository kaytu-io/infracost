package google

import (
	"github.com/kaytu-io/infracost/external/resources/google"
	"github.com/kaytu-io/infracost/external/schema"
)

func getComputeExternalVPNGatewayRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "google_compute_external_vpn_gateway",
		RFunc: NewComputeExternalVPNGateway,
	}
}
func NewComputeExternalVPNGateway(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &google.ComputeExternalVPNGateway{Address: d.Address, Region: d.Get("region").String()}
	r.PopulateUsage(u)
	return r.BuildResource()
}
