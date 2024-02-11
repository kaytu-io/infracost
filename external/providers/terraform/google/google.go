package google

import (
	"github.com/tidwall/gjson"

	"github.com/kaytu-io/infracost/external/providers/terraform/provider_schemas"
	"github.com/kaytu-io/infracost/external/schema"
)

var DefaultProviderRegion = "us-central1"

func GetDefaultRefIDFunc(d *schema.ResourceData) []string {

	defaultRefs := []string{d.Get("id").String()}

	if d.Get("self_link").Exists() {
		defaultRefs = append(defaultRefs, d.Get("self_link").String())
	}

	return defaultRefs
}

func DefaultCloudResourceIDFunc(d *schema.ResourceData) []string {
	return []string{}
}

func GetSpecialContext(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{}
}

func GetResourceRegion(resourceType string, v gjson.Result) string {
	if v.Get("region").Exists() && v.Get("region").String() != "" {
		return v.Get("region").String()
	}

	return ""
}

func ParseTags(r *schema.ResourceData) *map[string]string {
	_, supportsLabels := provider_schemas.GoogleLabelsSupport[r.Type]
	rLabels := r.Get("labels").Map()

	_, supportsUserLabels := provider_schemas.GoogleUserLabelsSupport[r.Type]
	rUserLabels := r.Get("user_labels").Map()

	_, supportsSettingsUserLabels := provider_schemas.GoogleSettingsUserLabelsSupport[r.Type]
	rSettingsUserLabels := r.Get("settings.0.user_labels").Map()

	if !supportsLabels && len(rLabels) == 0 &&
		!supportsUserLabels && len(rUserLabels) == 0 &&
		!supportsSettingsUserLabels && len(rSettingsUserLabels) == 0 {
		return nil
	}

	tags := make(map[string]string)
	for k, v := range rLabels {
		tags[k] = v.String()
	}
	for k, v := range rUserLabels {
		tags[k] = v.String()
	}
	for k, v := range rSettingsUserLabels {
		tags[k] = v.String()
	}
	return &tags
}