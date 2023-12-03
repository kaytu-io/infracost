package google

import (
	"github.com/kaytu-io/infracost/external/resources/google"
	"github.com/kaytu-io/infracost/external/schema"
)

func getLoggingFolderBucketConfigRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "google_logging_folder_bucket_config",
		RFunc: NewLoggingFolderBucketConfig,
	}
}

func NewLoggingFolderBucketConfig(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &google.Logging{
		Address: d.Address,
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
