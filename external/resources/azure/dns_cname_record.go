package azure

import (
	"github.com/kaytu-io/infracost/external/resources"
	"github.com/kaytu-io/infracost/external/schema"
)

type DNSCNameRecord struct {
	Address        string
	Region         string
	MonthlyQueries *int64 `infracost_usage:"monthly_queries"`
}

var DNSCNameRecordUsageSchema = []*schema.UsageItem{{Key: "monthly_queries", ValueType: schema.Int64, DefaultValue: 0}}

func (r *DNSCNameRecord) PopulateUsage(u *schema.UsageData) {
	resources.PopulateArgsWithUsage(r, u)
}

func (r *DNSCNameRecord) BuildResource() *schema.Resource {
	return &schema.Resource{
		Name:           r.Address,
		CostComponents: dnsQueriesCostComponent(r.Region, r.MonthlyQueries), UsageSchema: DNSCNameRecordUsageSchema,
	}
}
