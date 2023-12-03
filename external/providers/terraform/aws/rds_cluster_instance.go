package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getRDSClusterInstanceRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_rds_cluster_instance",
		RFunc: NewRDSClusterInstance,
	}
}

func NewRDSClusterInstance(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	piEnabled := d.Get("performance_insights_enabled").Bool()
	piLongTerm := piEnabled && d.Get("performance_insights_retention_period").Int() > 7

	r := &aws.RDSClusterInstance{
		Address:                              d.Address,
		Region:                               d.Get("region").String(),
		InstanceClass:                        d.Get("instance_class").String(),
		IOOptimized:                          false, // IO Optimized isn't supported by terraform yet
		Engine:                               d.Get("engine").String(),
		PerformanceInsightsEnabled:           piEnabled,
		PerformanceInsightsLongTermRetention: piLongTerm,
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
