package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getNeptuneClusterRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_neptune_cluster",
		RFunc: NewNeptuneCluster,
	}
}

func NewNeptuneCluster(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.NeptuneCluster{
		Address:               d.Address,
		Region:                d.Get("region").String(),
		BackupRetentionPeriod: d.Get("backup_retention_period").Int(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
