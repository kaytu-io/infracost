package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getSignalRServiceRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_signalr_service",
		CoreRFunc: newSignalRService,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}

func newSignalRService(d *schema.ResourceData) schema.CoreResource {
	region := lookupRegion(d, []string{"resource_group_name"})
	return &azure.SignalRService{
		Address:     d.Address,
		Region:      region,
		SkuName:     d.Get("sku.0.name").String(),
		SkuCapacity: d.Get("sku.0.capacity").Int(),
	}
}
