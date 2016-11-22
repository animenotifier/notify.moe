package main

import (
	"bytes"

	"github.com/aerogo/aero"
)

func init() {
	plusRoute := []byte("/+")
	plusRouteAjax := []byte("/_/+")

	// This will rewrite /+UserName requests to /user/UserName
	app.Rewrite(func(ctx *aero.RewriteContext) {
		requestURI := ctx.URIBytes()

		if bytes.HasPrefix(requestURI, plusRoute) {
			newURI := []byte("/user/")
			userName := requestURI[2:]
			newURI = append(newURI, userName...)
			ctx.SetURIBytes(newURI)
		} else if bytes.HasPrefix(requestURI, plusRouteAjax) {
			newURI := []byte("/_/user/")
			userName := requestURI[4:]
			newURI = append(newURI, userName...)
			ctx.SetURIBytes(newURI)
		}
	})
}
