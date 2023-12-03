package google

import (
	"github.com/kaytu-io/infracost/external/resources/google"
	"github.com/kaytu-io/infracost/external/schema"
)

func getComputeSnapshotRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:                "google_compute_snapshot",
		RFunc:               newComputeSnapshot,
		ReferenceAttributes: []string{"source_disk"},
	}
}

func newComputeSnapshot(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	region := d.Get("region").String()

	size := computeSnapshotDiskSize(d)

	r := &google.ComputeSnapshot{
		Address:  d.Address,
		Region:   region,
		DiskSize: size,
	}
	r.PopulateUsage(u)

	return r.BuildResource()
}
