package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getS3BucketAnalyticsConfigurationRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_s3_bucket_analytics_configuration",
		RFunc: NewS3BucketAnalyticsConfiguration,
	}
}

func NewS3BucketAnalyticsConfiguration(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.S3BucketAnalyticsConfiguration{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
