package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getAzureRMExpressRouteConnectionRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_express_route_connection",
		RFunc: newExpressRouteConnection,
	}
}

func newExpressRouteConnection(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	e := &azure.ExpressRouteConnection{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	return e.BuildResource()
}
