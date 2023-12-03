package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getDNSCNameRecordRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_dns_cname_record",
		RFunc: NewDNSCNameRecord,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewDNSCNameRecord(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &azure.DNSCNameRecord{Address: d.Address, Region: lookupRegion(d, []string{"resource_group_name"})}
	r.PopulateUsage(u)
	return r.BuildResource()
}
