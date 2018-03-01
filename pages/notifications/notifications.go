package notifications

import (
	"net/http"
	"sort"

	"github.com/animenotifier/arn"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxNotifications = 30

// All shows all notifications sent so far.
func All(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in", nil)
	}

	var viewUser *arn.User
	nick := ctx.Get("nick")

	if nick != "" {
		viewUser, _ = arn.GetUserByNick(nick)
	} else {
		viewUser = user
	}

	notifications := viewUser.Notifications().Notifications()

	// Sort by date
	sort.Slice(notifications, func(i, j int) bool {
		return notifications[i].Created > notifications[j].Created
	})

	// Limit results
	if len(notifications) > maxNotifications {
		notifications = notifications[:maxNotifications]
	}

	return ctx.HTML(components.Notifications(notifications, viewUser, user))
}
