package pages

import (
	"strings"

	"github.com/aerogo/aero"
)

// Rewrite will rewrite the path before routing happens.
func Rewrite(ctx aero.RewriteContext) {
	requestURI := ctx.Path()

	// User profiles
	if strings.HasPrefix(requestURI, "/+") {
		newURI := "/user/"
		userName := requestURI[2:]
		ctx.SetPath(newURI + userName)
		return
	}

	if strings.HasPrefix(requestURI, "/_/+") {
		newURI := "/_/user/"
		userName := requestURI[4:]
		ctx.SetPath(newURI + userName)
		return
	}

	// Analytics
	if requestURI == "/dark-flame-master" {
		ctx.SetPath("/api/new/analytics")
		return
	}
}
