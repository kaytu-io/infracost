package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getDocDBClusterSnapshotRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_docdb_cluster_snapshot",
		RFunc: NewDocDBClusterSnapshot,
	}

}
func NewDocDBClusterSnapshot(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.DocDBClusterSnapshot{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
