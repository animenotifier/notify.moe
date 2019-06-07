package middleware

import (
	"net/http"

	"github.com/aerogo/aero"
)

// HTTPSRedirect middleware redirects to HTTPS if needed.
func HTTPSRedirect(next aero.Handler) aero.Handler {
	return func(ctx aero.Context) error {
		request := ctx.Request()
		userAgent := request.Header("User-Agent")
		isBrowser := userAgent != ""

		if isBrowser && request.Scheme() != "https" {
			return ctx.Redirect(http.StatusPermanentRedirect, "https://"+request.Host()+request.Path())
		}

		// Handle the request
		return next(ctx)
	}
}
