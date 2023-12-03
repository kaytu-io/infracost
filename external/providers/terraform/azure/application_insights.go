package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getApplicationInsightsRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_application_insights",
		RFunc: NewApplicationInsights,
	}
}
func NewApplicationInsights(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &azure.ApplicationInsights{
		Address:         d.Address,
		Region:          lookupRegion(d, []string{}),
		RetentionInDays: d.Get("retention_in_days").Int(),
	}
	r.PopulateUsage(u)
	return r.BuildResource()
}
