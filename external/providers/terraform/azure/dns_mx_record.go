package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getDNSMXRecordRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_dns_mx_record",
		RFunc: NewDNSMXRecord,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewDNSMXRecord(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &azure.DNSMXRecord{Address: d.Address, Region: lookupRegion(d, []string{"resource_group_name"})}
	r.PopulateUsage(u)
	return r.BuildResource()
}
