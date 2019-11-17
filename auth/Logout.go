package auth

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/aerogo/log"
	"github.com/animenotifier/notify.moe/arn"
)

// Logout is called when the user clicks the logout button.
// It deletes the "userId" from the session.
func Logout(app *aero.Application, authLog *log.Log) {
	app.Get("/logout", func(ctx aero.Context) error {
		if ctx.HasSession() {
			user := arn.GetUserFromContext(ctx)

			if user != nil {
				authLog.Info("%s logged out | %s | %s | %s | %s", user.Nick, user.ID, ctx.IP(), user.Email, user.RealName())
			}

			ctx.Session().Delete("userId")
		}

		return ctx.Redirect(http.StatusTemporaryRedirect, "/")
	})
}
