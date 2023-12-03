package main_test

import (
	"github.com/kaytu-io/infracost/external/testutil"
	"testing"
)

func TestAuthLoginHelpFlag(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"auth", "login", "--help"}, nil)
}
