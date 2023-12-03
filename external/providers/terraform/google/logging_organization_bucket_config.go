package google

import (
	"github.com/kaytu-io/infracost/external/resources/google"
	"github.com/kaytu-io/infracost/external/schema"
)

func getLoggingOrganizationBucketConfigRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "google_logging_organization_bucket_config",
		RFunc: NewLoggingOrganizationBucketConfig,
	}
}

func NewLoggingOrganizationBucketConfig(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &google.Logging{
		Address: d.Address,
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
