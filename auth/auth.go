package auth

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/utils"
)

const newUserStartRoute = "/welcome"

// Install installs the authentication routes in the application.
func Install(app *aero.Application) {
	// Google
	InstallGoogleAuth(app)

	// Facebook
	InstallFacebookAuth(app)

	// Twitter
	InstallTwitterAuth(app)

	// Logout
	app.Get("/logout", func(ctx *aero.Context) string {
		if ctx.HasSession() {
			user := utils.GetUser(ctx)

			if user != nil {
				authLog.Info("%s logged out | %s | %s | %s | %s", user.Nick, user.ID, ctx.RealIP(), user.Email, user.RealName())
			}

			ctx.Session().Delete("userId")
		}

		return ctx.Redirect("/")
	})
}
