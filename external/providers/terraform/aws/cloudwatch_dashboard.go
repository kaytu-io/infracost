package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getCloudwatchDashboardRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_cloudwatch_dashboard",
		RFunc: NewCloudwatchDashboard,
	}
}
func NewCloudwatchDashboard(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.CloudwatchDashboard{
		Address: d.Address,
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
