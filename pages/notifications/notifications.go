package notifications

import (
	"net/http"
	"sort"

	"github.com/animenotifier/notify.moe/arn"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
)

const maxNotifications = 30

// ByUser shows all notifications sent to the given user.
func ByUser(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in")
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
