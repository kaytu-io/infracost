package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getAppServicePlanRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_app_service_plan",
		RFunc: NewAppServicePlan,
	}
}
func NewAppServicePlan(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &azure.AppServicePlan{
		Address:     d.Address,
		Region:      lookupRegion(d, []string{}),
		SKUSize:     d.Get("sku.0.size").String(),
		SKUCapacity: d.Get("sku.0.capacity").Int(),
		Kind:        d.Get("kind").String(),
	}
	r.PopulateUsage(u)
	return r.BuildResource()
}
