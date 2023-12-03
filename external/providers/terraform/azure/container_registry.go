package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getContainerRegistryRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_container_registry",
		RFunc: NewContainerRegistry,
	}
}
func NewContainerRegistry(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &azure.ContainerRegistry{
		Address:                 d.Address,
		Region:                  lookupRegion(d, []string{}),
		GeoReplicationLocations: len(d.Get("georeplications").Array()),
		SKU:                     d.Get("sku").String(),
	}
	r.PopulateUsage(u)
	return r.BuildResource()
}
