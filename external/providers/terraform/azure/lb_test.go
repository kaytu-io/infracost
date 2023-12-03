package azure_test

import (
	"testing"

	"github.com/infracost/infracost/external/providers/terraform/tftest"
)

func TestAzureRMLoadBalancerGoldenFile(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "lb_test")
}
