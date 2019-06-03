package main

import (
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Sending notifications")

	defer color.Green("Finished")
	defer arn.Node.Close()

	for user := range arn.StreamUsers() {
		user.SendNotification(&arn.PushNotification{
			Title:   "You have received a gift!",
			Message: "Akyoto gifted you the item \"PRO Account - 1 month\".",
			Icon:    "https://media.notify.moe/images/avatars/large/4J6qpK1ve.png?1545634334",
			Link:    "https://notify.moe/thread/Sw0WDbsig",
		})
	}
}
