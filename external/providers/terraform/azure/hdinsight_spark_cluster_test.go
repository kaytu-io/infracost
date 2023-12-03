package azure_test

import (
	"testing"

	"github.com/infracost/infracost/external/providers/terraform/tftest"
)

func TestAzureRMHDInsightSparkClusterGoldenFile(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "hdinsight_spark_cluster_test") //nolint:misspell
}
