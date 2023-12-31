package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getMonitorDataCollectionRuleRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_monitor_data_collection_rule",
		CoreRFunc: newMonitorDataCollectionRule,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}

func newMonitorDataCollectionRule(d *schema.ResourceData) schema.CoreResource {
	region := lookupRegion(d, []string{"resource_group_name"})
	return &azure.MonitorDataCollectionRule{
		Address: d.Address,
		Region:  region,
	}
}
