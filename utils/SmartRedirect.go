package utils

import (
	"net/http"
	"strings"

	"github.com/aerogo/aero"
)

// SmartRedirect automatically adds the /_ prefix to the URI if required.
func SmartRedirect(ctx aero.Context, uri string) error {
	// Redirect
	prefix := ""

	if strings.HasPrefix(ctx.Path(), "/_") {
		prefix = "/_"
	}

	return ctx.Redirect(http.StatusTemporaryRedirect, prefix+uri)
}
