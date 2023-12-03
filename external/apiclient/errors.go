package apiclient

import (
	"regexp"

	"github.com/kaytu-io/infracost/external/clierror"
	"github.com/kaytu-io/infracost/external/config"
	"github.com/kaytu-io/infracost/external/ui"
)

// Regex for finding potential URLs and file paths in error messages
// so we can sanitize them. This regex might be too greedy and match
// too many things, but it's a start.
var pathRegex = regexp.MustCompile(`(\w+:)?[\.~\w]*[\/\\]+([^\s:'\"\]]+)|([\w+-]\.\w{2,})`)

func ReportCLIError(ctx *config.RunContext, cliErr error, replacePath bool) error {
	if ctx.Config.IsSelfHosted() {
		return nil
	}

	errMsg := ui.StripColor(cliErr.Error())
	stacktrace := ""

	sanitizedErr, ok := cliErr.(clierror.SanitizedError)
	if ok {
		errMsg = ui.StripColor(sanitizedErr.SanitizedError())
		stacktrace = sanitizedErr.SanitizedStack()
	}

	if replacePath {
		errMsg = pathRegex.ReplaceAllString(errMsg, "REPLACED_PATH")
	}

	d := ctx.EventEnv()
	d["error"] = errMsg
	d["stacktrace"] = stacktrace

	c := GetPricingAPIClient(ctx)
	return c.AddEvent("infracost-error", d)
}
