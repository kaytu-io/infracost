package aws

import (
	"github.com/kaytu-io/infracost/external/resources"
	"github.com/kaytu-io/infracost/external/schema"
)

type KMSExternalKey struct {
	Address string
	Region  string
}

var KMSExternalKeyUsageSchema = []*schema.UsageItem{}

func (r *KMSExternalKey) PopulateUsage(u *schema.UsageData) {
	resources.PopulateArgsWithUsage(r, u)
}

func (r *KMSExternalKey) BuildResource() *schema.Resource {
	kmsKey := &KMSKey{
		Region: r.Region,
	}

	return &schema.Resource{
		Name:           r.Address,
		CostComponents: []*schema.CostComponent{kmsKey.customerMasterKeyCostComponent()},
		UsageSchema:    KMSExternalKeyUsageSchema,
	}
}
