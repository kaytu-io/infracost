package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getCloudwatchLogGroupItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_cloudwatch_log_group",
		RFunc: NewCloudwatchLogGroup,
	}
}
func NewCloudwatchLogGroup(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.CloudwatchLogGroup{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
