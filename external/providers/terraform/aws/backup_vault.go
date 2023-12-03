package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getBackupVaultRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_backup_vault",
		RFunc: NewBackupVault,
		Notes: []string{"AWS Storage Gateway Volume Backup prices could not be found in the AWS pricing data."},
	}
}
func NewBackupVault(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.BackupVault{Address: d.Address, Region: d.Get("region").String()}
	r.PopulateUsage(u)
	return r.BuildResource()
}
