package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getNetworkWatcherRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_network_watcher",
		CoreRFunc: newNetworkWatcher,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}

func newNetworkWatcher(d *schema.ResourceData) schema.CoreResource {
	region := lookupRegion(d, []string{"resource_group_name"})
	return &azure.NetworkWatcher{
		Address: d.Address,
		Region:  region,
	}
}
