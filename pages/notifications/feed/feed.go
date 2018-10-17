package notificationsfeed

import (
	"github.com/aerogo/aero"
)

// maxNotifications indicates how many notifications are shown in the feed.
const maxNotifications = 20

// RSS returns a notifications feed in RSS format.
func RSS(ctx *aero.Context) string {
	return "reserved"
}

// JSON returns a notifications feed in JSON format.
func JSON(ctx *aero.Context) string {
	return "reserved"
}

// Atom returns a notifications feed in Atom format.
func Atom(ctx *aero.Context) string {
	return "reserved"
}
