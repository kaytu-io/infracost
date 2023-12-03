package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getNewKMSExternalKeyRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_kms_external_key",
		RFunc: NewKMSExternalKey,
	}
}

func NewKMSExternalKey(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.KMSExternalKey{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
