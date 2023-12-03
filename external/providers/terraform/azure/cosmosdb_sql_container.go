package azure

import (
	"github.com/kaytu-io/infracost/external/schema"
)

func GetAzureRMCosmosdbSQLContainerRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_cosmosdb_sql_container",
		RFunc: NewAzureRMCosmosdb,
		ReferenceAttributes: []string{
			"account_name",
			"resource_group_name",
		},
	}
}
