package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getDNSPtrRecordRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_dns_ptr_record",
		RFunc: NewDNSPtrRecord,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewDNSPtrRecord(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &azure.DNSPtrRecord{Address: d.Address, Region: lookupRegion(d, []string{"resource_group_name"})}
	r.PopulateUsage(u)
	return r.BuildResource()
}
