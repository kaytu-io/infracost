package google

import (
	"github.com/kaytu-io/infracost/external/resources/google"
	"github.com/kaytu-io/infracost/external/schema"
)

func getComputeImageRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:                "google_compute_image",
		RFunc:               newComputeImage,
		ReferenceAttributes: []string{"source_disk", "source_image", "source_snapshot"},
	}
}

func newComputeImage(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	region := d.Get("region").String()

	storageSize := computeImageDiskSize(d)

	r := &google.ComputeImage{
		Address:     d.Address,
		Region:      region,
		StorageSize: storageSize,
	}
	r.PopulateUsage(u)

	return r.BuildResource()
}
