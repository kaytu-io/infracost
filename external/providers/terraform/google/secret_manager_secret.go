package google

import (
	"github.com/kaytu-io/infracost/external/resources/google"
	"github.com/kaytu-io/infracost/external/schema"

	"github.com/tidwall/gjson"
)

func getSecretManagerSecretRegistryItem() *schema.RegistryItem {
	rfunc := func(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {

		r := newSecretManagerSecret(d)
		r.PopulateUsage(u)

		return r.BuildResource()
	}

	return &schema.RegistryItem{
		Name:  "google_secret_manager_secret",
		RFunc: rfunc,
	}
}

func newSecretManagerSecret(d *schema.ResourceData) *google.SecretManagerSecret {
	replicasCount := 1

	replications := d.Get("replication.0.user_managed.0")
	if replications.Type != gjson.Null && len(replications.Array()) > 0 {
		replicasCount = len(replications.Get("replicas").Array())
	}

	return &google.SecretManagerSecret{
		Address:              d.Address,
		Region:               d.Get("region").String(),
		ReplicationLocations: int64(replicasCount),
	}
}
