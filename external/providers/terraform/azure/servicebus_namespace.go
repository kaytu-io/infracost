package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getServiceBusNamespaceRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_servicebus_namespace",
		CoreRFunc: newServiceBusNamespace,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}

func newServiceBusNamespace(d *schema.ResourceData) schema.CoreResource {
	region := lookupRegion(d, []string{"resource_group_name"})
	return &azure.ServiceBusNamespace{
		Address:  d.Address,
		Region:   region,
		SKU:      d.Get("sku").String(),
		Capacity: d.Get("capacity").Int(),
	}
}
