package google

import (
	"github.com/kaytu-io/infracost/external/resources/google"
	"github.com/kaytu-io/infracost/external/schema"
)

func getSecretManagerSecretVersionRegistryItem() *schema.RegistryItem {
	rfunc := func(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {

		r := newSecretManagerSecretVersion(d)
		r.PopulateUsage(u)

		return r.BuildResource()
	}

	return &schema.RegistryItem{
		Name:  "google_secret_manager_secret_version",
		RFunc: rfunc,
		ReferenceAttributes: []string{
			"secret",
		},
	}
}

func newSecretManagerSecretVersion(d *schema.ResourceData) *google.SecretManagerSecretVersion {
	replicasCount := int64(1)

	secretReferences := d.References("secret")
	if len(secretReferences) > 0 {
		secret := newSecretManagerSecret(secretReferences[0])
		replicasCount = secret.ReplicationLocations
	}

	return &google.SecretManagerSecretVersion{
		Address:              d.Address,
		Region:               d.Get("region").String(),
		ReplicationLocations: replicasCount,
	}
}
