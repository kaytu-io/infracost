package aws

import (
	"github.com/infracost/infracost/external/resources/aws"
	"github.com/infracost/infracost/external/schema"
)

func getRoute53ResolverEndpointRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_route53_resolver_endpoint",
		RFunc: NewRoute53ResolverEndpoint,
	}
}

func NewRoute53ResolverEndpoint(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.Route53ResolverEndpoint{
		Address:           d.Address,
		Region:            d.Get("region").String(),
		ResolverEndpoints: int64(len(d.Get("ip_address").Array())),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
