package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getAppServiceEnvironmentRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_app_service_environment",
		RFunc: NewAppServiceEnvironment,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewAppServiceEnvironment(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &azure.AppServiceEnvironment{
		Address:     d.Address,
		Region:      lookupRegion(d, []string{"resource_group_name"}),
		PricingTier: d.Get("pricing_tier").String(),
	}
	r.PopulateUsage(u)
	return r.BuildResource()
}
