package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getDNSARecordRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_dns_a_record",
		RFunc: NewDNSARecord,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewDNSARecord(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &azure.DNSARecord{Address: d.Address, Region: lookupRegion(d, []string{"resource_group_name"})}
	r.PopulateUsage(u)
	return r.BuildResource()
}
