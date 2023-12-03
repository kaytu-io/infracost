package azure

import (
	"fmt"

	"github.com/kaytu-io/infracost/external/resources"
	"github.com/kaytu-io/infracost/external/schema"

	"strings"

	"github.com/shopspring/decimal"
)

type ApplicationInsightsWebTest struct {
	Address string
	Region  string
	Kind    string
	Enabled bool
}

var ApplicationInsightsWebTestUsageSchema = []*schema.UsageItem{}

func (r *ApplicationInsightsWebTest) PopulateUsage(u *schema.UsageData) {
	resources.PopulateArgsWithUsage(r, u)
}

func (r *ApplicationInsightsWebTest) BuildResource() *schema.Resource {
	costComponents := []*schema.CostComponent{}

	if r.Kind != "" {
		if strings.ToLower(r.Kind) == "multistep" && r.Enabled {
			costComponents = append(costComponents, &schema.CostComponent{
				Name:            "Multi-step web test",
				Unit:            "test",
				UnitMultiplier:  decimal.NewFromInt(1),
				MonthlyQuantity: decimalPtr(decimal.NewFromInt(1)),
				ProductFilter: &schema.ProductFilter{
					VendorName:    strPtr("azurerm"),
					Region:        strPtr(r.Region),
					Service:       strPtr("Application Insights"),
					ProductFamily: strPtr("Management and Governance"),
					AttributeFilters: []*schema.AttributeFilter{
						{Key: "meterName", ValueRegex: strPtr(fmt.Sprintf("/^%s$/i", "Multi-step Web Test"))},
						{Key: "skuName", ValueRegex: strPtr(fmt.Sprintf("/^%s$/i", "Enterprise"))},
					},
				},
			})
		}
	}

	if len(costComponents) == 0 {
		return &schema.Resource{
			Name:      r.Address,
			IsSkipped: true,
			NoPrice:   true, UsageSchema: ApplicationInsightsWebTestUsageSchema,
		}
	}

	return &schema.Resource{
		Name:           r.Address,
		CostComponents: costComponents, UsageSchema: ApplicationInsightsWebTestUsageSchema,
	}

}
