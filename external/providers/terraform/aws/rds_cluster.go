package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getRDSClusterRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_rds_cluster",
		RFunc: NewRDSCluster,
	}
}

func NewRDSCluster(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.RDSCluster{
		Address:               d.Address,
		Region:                d.Get("region").String(),
		Engine:                d.GetStringOrDefault("engine", "aurora"),
		BackupRetentionPeriod: d.GetInt64OrDefault("backup_retention_period", 1),
		EngineMode:            d.GetStringOrDefault("engine_mode", "provisioned"),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
