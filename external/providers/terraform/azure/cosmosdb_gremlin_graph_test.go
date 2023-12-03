package azure_test

import (
	"testing"

	"github.com/kaytu-io/infracost/external/providers/terraform/tftest"
)

func TestAzureRMCosmosDBGremlinGraphGoldenFile(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "cosmosdb_gremlin_graph_test")
}
