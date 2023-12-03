package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getMonitorMetricAlertRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_monitor_metric_alert",
		CoreRFunc: newMonitorMetricAlert,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}

func newMonitorMetricAlert(d *schema.ResourceData) schema.CoreResource {
	region := lookupRegion(d, []string{"resource_group_name"})

	scopeCount := 1 // default scope is the azurerm subscription, so count == 1
	if !d.IsEmpty("scopes") {

		scopeCount = len(d.Get("scopes").Array())
	}

	criteriaDimensionsCount := 0
	for _, c := range d.Get("criteria").Array() {
		criteriaDimensionsCount += len(c.Get("dimension").Array())
	}

	dynamicCriteriaDimensionsCount := 0
	for _, c := range d.Get("dynamic_criteria").Array() {
		dynamicCriteriaDimensionsCount += len(c.Get("dimension").Array())
	}

	return &azure.MonitorMetricAlert{
		Address:                        d.Address,
		Region:                         region,
		Enabled:                        d.GetBoolOrDefault("enabled", true),
		ScopeCount:                     scopeCount,
		CriteriaDimensionsCount:        criteriaDimensionsCount,
		DynamicCriteriaDimensionsCount: dynamicCriteriaDimensionsCount,
	}
}
