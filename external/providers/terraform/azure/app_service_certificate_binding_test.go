package azure_test

import (
	"testing"

	"github.com/kaytu-io/infracost/external/providers/terraform/tftest"
)

func TestAzureRMAppServiceCertificateBinding(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "app_service_certificate_binding_test")
}
