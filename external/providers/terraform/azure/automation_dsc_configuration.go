package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getAutomationDSCConfigurationRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_automation_dsc_configuration",
		RFunc: NewAutomationDSCConfiguration,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewAutomationDSCConfiguration(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &azure.AutomationDSCConfiguration{Address: d.Address, Region: lookupRegion(d, []string{"resource_group_name"})}
	r.PopulateUsage(u)
	return r.BuildResource()
}
