package azure_test

import (
	"testing"

	"github.com/infracost/infracost/external/providers/terraform/tftest"
)

func TestEventgridSystemTopic(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "eventgrid_system_topic_test")
}
