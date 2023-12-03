package aws_test

import (
	"testing"

	"github.com/infracost/infracost/external/providers/terraform/tftest"
)

func TestKMSKey(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "kms_key_test")
}
