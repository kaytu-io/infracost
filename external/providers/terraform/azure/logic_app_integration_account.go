package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getLogicAppIntegrationAccountRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_logic_app_integration_account",
		CoreRFunc: newLogicAppIntegrationAccount,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}

func newLogicAppIntegrationAccount(d *schema.ResourceData) schema.CoreResource {
	region := lookupRegion(d, []string{"resource_group_name"})

	return azure.NewLogicAppIntegrationAccount(d.Address, region, d.GetStringOrDefault("sku_name", "free"))
}
