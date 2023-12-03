package google

import (
	"github.com/kaytu-io/infracost/external/resources/google"
	"github.com/kaytu-io/infracost/external/schema"
)

func getKMSCryptoKeyRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "google_kms_crypto_key",
		RFunc: NewKMSCryptoKey,
	}
}
func NewKMSCryptoKey(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &google.KMSCryptoKey{Address: d.Address, Region: d.Get("region").String(), Algorithm: d.Get("version_template.0.algorithm").String(), ProtectionLevel: d.Get("version_template.0.protection_level").String(), RotationPeriod: d.Get("rotation_period").String(), VersionTemplate: d.Get("version_template").String()}
	r.PopulateUsage(u)
	return r.BuildResource()
}
