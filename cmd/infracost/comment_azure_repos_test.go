package main_test

import (
	"testing"

	"github.com/infracost/infracost/external/testutil"
)

func TestCommentAzureReposHelp(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"comment", "azurerm-repos", "--help"}, nil)
}

func TestCommentAzureReposPullRequest(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(),
		[]string{"comment", "azurerm-repos", "--azurerm-access-token", "abc", "--repo-url", "https://dev.azure.com/my-org/my-project/_git/my-azure-repo", "--pull-request", "5", "--path", "./testdata/terraform_v0.14_breakdown.json", "--dry-run"},
		nil)
}
