package google

import (
	"github.com/kaytu-io/infracost/external/resources/google"
	"github.com/kaytu-io/infracost/external/schema"
)

func getPubSubSubscriptionRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "google_pubsub_subscription",
		RFunc: NewPubSubSubscription,
	}
}

func NewPubSubSubscription(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &google.PubSubSubscription{
		Address: d.Address,
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
