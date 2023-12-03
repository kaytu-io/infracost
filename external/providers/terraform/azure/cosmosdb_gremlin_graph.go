package azure

import (
	"github.com/kaytu-io/infracost/external/schema"
)

func GetAzureRMCosmosdbGremlinGraphRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_cosmosdb_gremlin_graph",
		RFunc: NewAzureRMCosmosdb,
		ReferenceAttributes: []string{
			"account_name",
			"resource_group_name",
		},
	}
}
