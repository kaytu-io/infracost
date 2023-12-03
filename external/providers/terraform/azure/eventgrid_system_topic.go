package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getEventgridSystemTopicRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name: "azurerm_eventgrid_system_topic",
		CoreRFunc: func(d *schema.ResourceData) schema.CoreResource {
			return &azure.EventGridTopic{
				Address: d.Address,
				Region:  lookupRegion(d, []string{"resource_group_name"}),
			}
		},
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
