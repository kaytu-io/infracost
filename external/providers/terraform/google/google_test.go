package google_test

import (
	"os"
	"testing"

	"github.com/kaytu-io/infracost/external/providers/terraform/tftest"
)

func TestMain(m *testing.M) {
	tftest.EnsurePluginsInstalled()
	code := m.Run()
	os.Exit(code)
}
