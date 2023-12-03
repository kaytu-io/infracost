package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getActiveDirectoryDomainServiceRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_active_directory_domain_service",
		RFunc: NewActiveDirectoryDomainService,
	}
}
func NewActiveDirectoryDomainService(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &azure.ActiveDirectoryDomainService{Address: d.Address, Region: lookupRegion(d, []string{}), SKU: d.Get("sku").String()}
	r.PopulateUsage(u)
	return r.BuildResource()
}
