package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getDNSSrvRecordRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_dns_srv_record",
		RFunc: NewDNSSrvRecord,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewDNSSrvRecord(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &azure.DNSSrvRecord{Address: d.Address, Region: lookupRegion(d, []string{"resource_group_name"})}
	r.PopulateUsage(u)
	return r.BuildResource()
}
