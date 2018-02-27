package notifications

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/utils"
)

// Test ...
func Test(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in", nil)
	}

	user.SendNotification(&arn.PushNotification{
		Title:   "Anime Notifier",
		Message: "Yay, it works!",
		Icon:    "https://" + ctx.App.Config.Domain + "/images/brand/220.png",
	})

	return "ok"
}
