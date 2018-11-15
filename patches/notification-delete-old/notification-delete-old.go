package main

import (
	"fmt"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

const maxNotificationsPerUser = 80

func main() {
	color.Yellow("Deleting old notifications")

	defer color.Green("Finished")
	defer arn.Node.Close()

	for user := range arn.StreamUsers() {
		notificationCount := len(user.Notifications().Items)

		if notificationCount > maxNotificationsPerUser {
			fmt.Println(user.Nick, notificationCount)
		}
	}
}
