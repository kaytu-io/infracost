package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getAutomationJobScheduleRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_automation_job_schedule",
		RFunc: NewAutomationJobSchedule,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewAutomationJobSchedule(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &azure.AutomationJobSchedule{Address: d.Address, Region: lookupRegion(d, []string{"resource_group_name"})}
	r.PopulateUsage(u)
	return r.BuildResource()
}
