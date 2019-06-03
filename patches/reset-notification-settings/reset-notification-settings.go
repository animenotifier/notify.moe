package main

import "github.com/animenotifier/notify.moe/arn"

func main() {
	defer arn.Node.Close()

	for user := range arn.StreamUsers() {
		settings := user.Settings()
		defaultSettings := arn.DefaultNotificationSettings()

		settings.Notification.QuoteLikes = defaultSettings.QuoteLikes
		settings.Notification.SoundTrackLikes = defaultSettings.SoundTrackLikes
		settings.Notification.GroupPostLikes = defaultSettings.GroupPostLikes
		settings.Notification.ForumLikes = defaultSettings.ForumLikes
		settings.Notification.NewFollowers = defaultSettings.NewFollowers
		settings.Save()
	}
}
