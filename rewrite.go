package main

import (
	"strings"

	"github.com/aerogo/aero"
)

func init() {
	plusRoute := "/+"
	plusRouteAjax := "/_/+"

	// This will rewrite /+UserName requests to /user/UserName
	app.Rewrite(func(ctx *aero.RewriteContext) {
		requestURI := ctx.URI()

		if strings.HasPrefix(requestURI, plusRoute) {
			newURI := "/user/"
			userName := requestURI[2:]
			ctx.SetURI(newURI + userName)
		} else if strings.HasPrefix(requestURI, plusRouteAjax) {
			newURI := "/_/user/"
			userName := requestURI[4:]
			ctx.SetURI(newURI + userName)
		}
	})
}
