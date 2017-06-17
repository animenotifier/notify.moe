package auth

import "github.com/aerogo/aero"

// Install ...
func Install(app *aero.Application) {
	// Google
	InstallGoogleAuth(app)

	// Logout
	app.Get("/logout", func(ctx *aero.Context) string {
		if ctx.HasSession() {
			ctx.Session().Set("userId", nil)
		}

		return ctx.Redirect("/")
	})
}
