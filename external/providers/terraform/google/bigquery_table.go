package google

import (
	"github.com/infracost/infracost/external/resources/google"
	"github.com/infracost/infracost/external/schema"
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
