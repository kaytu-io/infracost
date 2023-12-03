package azure_test

import (
	"testing"

	"github.com/infracost/infracost/external/providers/terraform/tftest"
)

func TestLogicAppIntegrationAccount(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "logic_app_integration_account_test")
}
