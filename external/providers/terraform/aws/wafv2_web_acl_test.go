package aws_test

import (
	"testing"

	"github.com/kaytu-io/infracost/external/providers/terraform/tftest"
)

func TestWAFv2WebACLGoldenFile(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "wafv2_web_acl_test")
}
