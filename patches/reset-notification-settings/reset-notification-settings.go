package main

import "github.com/animenotifier/arn"

func main() {
	defer arn.Node.Close()

	for user := range arn.StreamUsers() {
		settings := user.Settings()
		settings.Notification = arn.DefaultNotificationSettings()
		settings.Save()
	}
}
