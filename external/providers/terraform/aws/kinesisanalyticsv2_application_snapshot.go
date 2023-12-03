package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getKinesisAnalyticsV2ApplicationSnapshotRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_kinesisanalyticsv2_application_snapshot",
		RFunc: NewKinesisAnalyticsV2ApplicationSnapshot,
	}
}

func NewKinesisAnalyticsV2ApplicationSnapshot(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.KinesisAnalyticsV2ApplicationSnapshot{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
