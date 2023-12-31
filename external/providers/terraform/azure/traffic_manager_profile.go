package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
	"strings"
)

func getTrafficManagerProfileRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_traffic_manager_profile",
		CoreRFunc: newTrafficManagerProfile,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}

func newTrafficManagerProfile(d *schema.ResourceData) schema.CoreResource {
	region := lookupRegion(d, []string{"resource_group_name"})

	return &azure.TrafficManagerProfile{
		Address:            d.Address,
		Region:             region,
		Enabled:            trafficManagerProfileEnabled(d),
		TrafficViewEnabled: d.Get("traffic_view_enabled").Bool(),
	}
}

func trafficManagerProfileEnabled(d *schema.ResourceData) bool {
	return strings.ToLower(d.GetStringOrDefault("profile_status", "enabled")) == "enabled"
}
