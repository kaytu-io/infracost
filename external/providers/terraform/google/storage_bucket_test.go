package google_test

import (
	"testing"

	"github.com/kaytu-io/infracost/external/providers/terraform/tftest"
)

func TestStorageBucket(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "storage_bucket_test")
}
