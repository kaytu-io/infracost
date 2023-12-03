package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getAPIManagementRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_api_management",
		RFunc: NewAPIManagement,
		ReferenceAttributes: []string{
			"certificate_id",
		},
	}
}
func NewAPIManagement(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &azure.APIManagement{Address: d.Address, SKUName: d.Get("sku_name").String(), Region: lookupRegion(d, []string{})}
	r.PopulateUsage(u)
	return r.BuildResource()
}
