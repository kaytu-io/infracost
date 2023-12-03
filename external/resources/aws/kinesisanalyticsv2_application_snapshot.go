package aws

import (
	"github.com/kaytu-io/infracost/external/resources"
	"github.com/kaytu-io/infracost/external/schema"
	"github.com/shopspring/decimal"
)

type KinesisAnalyticsV2ApplicationSnapshot struct {
	Address                    string
	Region                     string
	DurableApplicationBackupGB *float64 `infracost_usage:"durable_application_backup_gb"`
}

var KinesisAnalyticsV2ApplicationSnapshotUsageSchema = []*schema.UsageItem{
	{Key: "durable_application_backup_gb", ValueType: schema.Float64, DefaultValue: 0},
}

func (r *KinesisAnalyticsV2ApplicationSnapshot) PopulateUsage(u *schema.UsageData) {
	resources.PopulateArgsWithUsage(r, u)
}

func (r *KinesisAnalyticsV2ApplicationSnapshot) BuildResource() *schema.Resource {
	var durableApplicationBackupGB *decimal.Decimal
	if r.DurableApplicationBackupGB != nil {
		durableApplicationBackupGB = decimalPtr(decimal.NewFromFloat(*r.DurableApplicationBackupGB))
	}

	v2App := &KinesisAnalyticsV2Application{
		Region:                     r.Region,
		DurableApplicationBackupGB: r.DurableApplicationBackupGB,
	}

	return &schema.Resource{
		Name:           r.Address,
		CostComponents: []*schema.CostComponent{v2App.backupCostComponent(durableApplicationBackupGB)},
		UsageSchema:    KinesisAnalyticsV2ApplicationSnapshotUsageSchema,
	}
}
