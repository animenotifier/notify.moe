package utils

import (
	"strings"

	"github.com/aerogo/aero"
)

// SmartRedirect automatically adds the /_ prefix to the URI if required.
func SmartRedirect(ctx *aero.Context, uri string) string {
	// Redirect
	prefix := ""

	if strings.HasPrefix(ctx.URI(), "/_") {
		prefix = "/_"
	}

	return ctx.Redirect(prefix + uri)
}
