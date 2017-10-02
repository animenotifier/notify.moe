package main

import (
	"strings"

	"github.com/aerogo/aero"
)

// rewrite will rewrite certain routes
func rewrite(ctx *aero.RewriteContext) {
	requestURI := ctx.URI()

	// User profiles
	if strings.HasPrefix(requestURI, "/+") {
		newURI := "/user/"
		userName := requestURI[2:]
		ctx.SetURI(newURI + userName)
		return
	}

	if strings.HasPrefix(requestURI, "/_/+") {
		newURI := "/_/user/"
		userName := requestURI[4:]
		ctx.SetURI(newURI + userName)
		return
	}

	// Search
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

	// Analytics
	if requestURI == "/dark-flame-master" {
		ctx.SetURI("/api/new/analytics")
		return
	}
}
