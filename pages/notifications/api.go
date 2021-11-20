package notifications

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/aerogo/aero"
	"github.com/aerogo/aero/event"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
)

// CountUnseen sends the number of unseen notifications.
func CountUnseen(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in")
	}

	unseen := user.Notifications().CountUnseen()

	return ctx.Text(strconv.Itoa(unseen))
}

// MarkNotificationsAsSeen marks all notifications as seen.
func MarkNotificationsAsSeen(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in")
	}

	notifications := user.Notifications().Notifications()

	for _, notification := range notifications {
		notification.Seen = arn.DateTimeUTC()
		notification.Save()
	}

	// Update the counter on all clients
	user.BroadcastEvent(event.New("notificationCount", 0))

	return nil
}

// Latest returns the latest notifications.
func Latest(ctx aero.Context) error {
	userID := ctx.Get("id")
	user, err := arn.GetUser(userID)

	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Invalid user ID")
	}

	notifications := user.Notifications().Notifications()

	// Sort by date
	sort.Slice(notifications, func(i, j int) bool {
		return notifications[i].Created > notifications[j].Created
	})

	if len(notifications) > maxNotifications {
		notifications = notifications[:maxNotifications]
	}

	return ctx.JSON(notifications)
}

// Test sends a test notification to the logged in user.
func Test(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in")
	}

	user.SendNotification(&arn.PushNotification{
		Title:   "Anime Notifier",
		Message: "Yay, it works!",
		Icon:    "https://" + assets.Domain + "/images/brand/220.png",
		Type:    arn.NotificationTypeTest,
	})

	return nil
}
