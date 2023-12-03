package azure

import (
	"github.com/kaytu-io/infracost/external/resources/azure"
	"github.com/kaytu-io/infracost/external/schema"
)

func getAppServiceCertificateOrderRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_app_service_certificate_order",
		RFunc: NewAppServiceCertificateOrder,
	}
}
func NewAppServiceCertificateOrder(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &azure.AppServiceCertificateOrder{Address: d.Address, ProductType: d.Get("product_type").String()}
	r.PopulateUsage(u)
	return r.BuildResource()
}
