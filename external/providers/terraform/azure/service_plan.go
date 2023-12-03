package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getServicePlanRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name: "azurerm_service_plan",
		CoreRFunc: func(d *schema.ResourceData) schema.CoreResource {
			return &azure.ServicePlan{
				Address:     d.Address,
				Region:      lookupRegion(d, []string{}),
				SKUName:     d.Get("sku_name").String(),
				WorkerCount: d.GetInt64OrDefault("worker_count", 1),
				OSType:      d.Get("os_type").String(),
			}
		},
	}
}
