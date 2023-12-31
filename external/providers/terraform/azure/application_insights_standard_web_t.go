package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getApplicationInsightsStandardWebTestRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_application_insights_standard_web_test",
		CoreRFunc: newApplicationInsightsStandardWebTest,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}

func newApplicationInsightsStandardWebTest(d *schema.ResourceData) schema.CoreResource {
	region := lookupRegion(d, []string{"resource_group_name"})
	return &azure.ApplicationInsightsStandardWebTest{
		Address:   d.Address,
		Region:    region,
		Enabled:   d.GetBoolOrDefault("enabled", true),
		Frequency: d.GetInt64OrDefault("frequency", 300),
	}
}
