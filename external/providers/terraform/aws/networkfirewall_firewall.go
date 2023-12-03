package aws

import (
	"github.com/infracost/infracost/external/resources/aws"
	"github.com/infracost/infracost/external/schema"
)

func getNetworkfirewallFirewallRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_networkfirewall_firewall",
		RFunc: newNetworkfirewallFirewall,
	}
}

func newNetworkfirewallFirewall(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	region := d.Get("region").String()
	r := &aws.NetworkfirewallFirewall{
		Address: d.Address,
		Region:  region,
	}
	r.PopulateUsage(u)

	return r.BuildResource()
}
