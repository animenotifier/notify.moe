package main

import (
	"bytes"

	"github.com/valyala/fasthttp"
)

func init() {
	plusRoute := []byte("/+")
	plusRouteAjax := []byte("/_/+")

	// This will rewrite /+UserName requests to /user/UserName
	app.Rewrite(func(ctx *fasthttp.RequestCtx) {
		requestURI := ctx.RequestURI()

		if bytes.HasPrefix(requestURI, plusRoute) {
			newURI := []byte("/user/")
			userName := requestURI[2:]
			newURI = append(newURI, userName...)
			ctx.Request.SetRequestURIBytes(newURI)
		} else if bytes.HasPrefix(requestURI, plusRouteAjax) {
			newURI := []byte("/_/user/")
			userName := requestURI[4:]
			newURI = append(newURI, userName...)
			ctx.Request.SetRequestURIBytes(newURI)
		}
	})
}
