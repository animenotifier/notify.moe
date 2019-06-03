package notifications

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxAllNotifications = 150

// All shows all notifications.
func All(ctx aero.Context) error {
	notifications, err := arn.AllNotifications()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not retrieve notification list", err)
	}

	// Sort by date
	sort.Slice(notifications, func(i, j int) bool {
		return notifications[i].Created > notifications[j].Created
	})

	// Limit results
	if len(notifications) > maxAllNotifications {
		notifications = notifications[:maxAllNotifications]
	}

	return ctx.HTML(components.AllNotifications(notifications))
}
