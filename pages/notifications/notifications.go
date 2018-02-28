package notifications

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// All shows all notifications sent so far.
func All(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in", nil)
	}

	notifications := user.Notifications().Notifications()

	// Sort by date
	sort.Slice(notifications, func(i, j int) bool {
		return notifications[i].Created > notifications[j].Created
	})

	return ctx.HTML(components.Notifications(notifications, user))
}

// Test sends a test notification to the logged in user.
func Test(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in", nil)
	}

	user.SendNotification(&arn.PushNotification{
		Title:   "Anime Notifier",
		Message: "Yay, it works!",
		Icon:    "https://" + ctx.App.Config.Domain + "/images/brand/220.png",
		Type:    arn.NotificationTypeTest,
	})

	return "ok"
}
