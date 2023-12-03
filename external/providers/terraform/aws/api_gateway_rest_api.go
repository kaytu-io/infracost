package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getAPIGatewayRestAPIRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_api_gateway_rest_api",
		RFunc: NewAPIGatewayRestAPI,
	}
}
func NewAPIGatewayRestAPI(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.APIGatewayRestAPI{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
