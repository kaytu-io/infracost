package azure

import (
	"github.com/kaytu-io/infracost/external/resources"
	"github.com/kaytu-io/infracost/external/schema"
)

type PrivateDNSTXTRecord struct {
	Address        string
	Region         string
	MonthlyQueries *int64 `infracost_usage:"monthly_queries"`
}

var PrivateDNSTXTRecordUsageSchema = []*schema.UsageItem{{Key: "monthly_queries", ValueType: schema.Int64, DefaultValue: 0}}

func (r *PrivateDNSTXTRecord) PopulateUsage(u *schema.UsageData) {
	resources.PopulateArgsWithUsage(r, u)
}

func (r *PrivateDNSTXTRecord) BuildResource() *schema.Resource {
	return &schema.Resource{
		Name:           r.Address,
		CostComponents: dnsQueriesCostComponent(r.Region, r.MonthlyQueries), UsageSchema: PrivateDNSTXTRecordUsageSchema,
	}
}
