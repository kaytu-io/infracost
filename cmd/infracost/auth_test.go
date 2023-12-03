package main_test

import (
	"github.com/kaytu-io/infracost/external/testutil"
	"testing"
)

func TestAuthNoArgs(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"auth"}, nil)
}

func TestAuthHelpFlag(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"auth", "--help"}, nil)
}
