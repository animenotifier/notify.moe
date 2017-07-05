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
			return
		}

		if strings.HasPrefix(requestURI, plusRouteAjax) {
			newURI := "/_/user/"
			userName := requestURI[4:]
			ctx.SetURI(newURI + userName)
			return
		}

		if strings.HasPrefix(requestURI, "/search/") {
			searchTerm := requestURI[len("/search/"):]
			ctx.Request.URL.RawQuery = "q=" + searchTerm
			ctx.SetURI("/search")
			return
		}

		if strings.HasPrefix(requestURI, "/_/search/") {
			searchTerm := requestURI[len("/_/search/"):]
			ctx.Request.URL.RawQuery = "q=" + searchTerm
			ctx.SetURI("/_/search")
			return
		}

		if requestURI == "/dark-flame-master" {
			ctx.SetURI("/api/analytics/new")
			return
		}
	})
}
