package azure

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/kaytu-io/infracost/external/schema"
)

const (
	vendorName = "azurerm"
)

var (
	priceFilterConsumption = &schema.PriceFilter{
		PurchaseOption: strPtr("Consumption"),
	}
)

func strPtr(s string) *string {
	return &s
}

func decimalPtr(d decimal.Decimal) *decimal.Decimal {
	return &d
}

func int32Ptr(i int32) *int32 {
	return &i
}

func floatPtrToDecimalPtr(f *float64) *decimal.Decimal {
	if f == nil {
		return nil
	}
	return decimalPtr(decimal.NewFromFloat(*f))
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func regexPtr(regex string) *string {
	return strPtr(fmt.Sprintf("/%s/i", regex))
}
