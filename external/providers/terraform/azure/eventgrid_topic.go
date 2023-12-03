package azure

import (
	"github.com/infracost/infracost/external/resources/azure"
	"github.com/infracost/infracost/external/schema"
)

func getEventgridTopicRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name: "azurerm_eventgrid_topic",
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
