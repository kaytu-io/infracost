package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getGlueCrawlerRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_glue_crawler",
		RFunc: newGlueCrawler,
	}
}

func newGlueCrawler(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	region := d.Get("region").String()
	r := &aws.GlueCrawler{
		Address: d.Address,
		Region:  region,
	}
	r.PopulateUsage(u)

	return r.BuildResource()
}
