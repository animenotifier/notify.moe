package main

import (
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

// HTMLEmailRenderer uses pixy templates to render the HTML for our emails.
type HTMLEmailRenderer struct{}

// Notification renders a notification email.
func (writer *HTMLEmailRenderer) Notification(notification *arn.Notification) string {
	return components.NotificationEmail(notification)
}
