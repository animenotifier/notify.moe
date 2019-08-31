package htmlemail

import (
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Renderer uses pixy templates to render the HTML for our emails.
type Renderer struct{}

// Notification renders a notification email.
func (writer *Renderer) Notification(notification *arn.Notification) string {
	return components.NotificationEmail(notification)
}
