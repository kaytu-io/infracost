package azure_test

import (
	"testing"

	"github.com/kaytu-io/infracost/external/providers/terraform/tftest"
)

func TestSQLElasticPool(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	opts := tftest.DefaultGoldenFileOptions()
	opts.CaptureLogs = true
	tftest.GoldenFileResourceTestsWithOpts(t, "sql_elasticpool_test", opts)
}
