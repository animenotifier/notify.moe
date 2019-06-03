package main

import (
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

const maxNotificationsPerUser = 30

func main() {
	color.Yellow("Deleting old notifications")

	defer color.Green("Finished")
	defer arn.Node.Close()

	count := 0

	for user := range arn.StreamUsers() {
		notifications := user.Notifications()
		notificationCount := len(notifications.Items)

		if notificationCount > maxNotificationsPerUser {
			cut := len(notifications.Items) - maxNotificationsPerUser
			deletedItems := notifications.Items[:cut]
			newItems := notifications.Items[cut:]

			for _, notificationID := range deletedItems {
				arn.DB.Delete("Notification", notificationID)
			}

			notifications.Items = newItems
			notifications.Save()

			count += len(deletedItems)
		}
	}

	color.Green("Deleted %d notifications", count)
}
