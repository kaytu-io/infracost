package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getAPIGatewayStageRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_api_gateway_stage",
		RFunc: NewAPIGatewayStage,
	}
}
func NewAPIGatewayStage(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.APIGatewayStage{
		Address:          d.Address,
		Region:           d.Get("region").String(),
		CacheClusterSize: d.Get("cache_cluster_size").Float(),
		CacheEnabled:     d.GetBoolOrDefault("cache_cluster_enabled", false),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
