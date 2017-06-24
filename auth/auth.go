package auth

import "github.com/aerogo/aero"
import "github.com/animenotifier/notify.moe/utils"

// Install ...
func Install(app *aero.Application) {
	// Google
	InstallGoogleAuth(app)

	// Logout
	app.Get("/logout", func(ctx *aero.Context) string {
		if ctx.HasSession() {
			user := utils.GetUser(ctx)

			if user != nil {
				authLog.Info("User logged out", user.ID, ctx.RealIP(), user.Email, user.RealName())
			}

			ctx.Session().Set("userId", nil)
		}

		return ctx.Redirect("/")
	})
}
