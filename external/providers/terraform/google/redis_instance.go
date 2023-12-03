package google

import (
	"github.com/kaytu-io/infracost/external/resources/google"
	"github.com/kaytu-io/infracost/external/schema"
)

func getRedisInstanceRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "google_redis_instance",
		RFunc: NewRedisInstance,
	}
}

func NewRedisInstance(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &google.RedisInstance{
		Address:      d.Address,
		Region:       d.Get("region").String(),
		MemorySizeGB: d.Get("memory_size_gb").Float(),
		Tier:         d.Get("tier").String(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
