package google

import (
	"github.com/kaytu-io/infracost/external/resources/google"
	"github.com/kaytu-io/infracost/external/schema"
)

func getStorageBucketRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:                "google_storage_bucket",
		RFunc:               NewStorageBucket,
		ReferenceAttributes: []string{},
	}
}
func NewStorageBucket(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &google.StorageBucket{
		Address:      d.Address,
		Region:       d.Get("region").String(),
		Location:     d.Get("location").String(),
		StorageClass: d.Get("storage_class").String(),
	}
	r.PopulateUsage(u)
	return r.BuildResource()
}
