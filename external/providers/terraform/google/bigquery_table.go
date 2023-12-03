package google

import (
	"github.com/kaytu-io/infracost/external/resources/google"
	"github.com/kaytu-io/infracost/external/schema"
)

func getBigQueryTableRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "google_bigquery_table",
		RFunc: NewBigQueryTable,
	}
}

func NewBigQueryTable(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &google.BigQueryTable{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
