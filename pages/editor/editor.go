package editor

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get ...
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil || (user.Role != "admin" && user.Role != "editor") {
		return ctx.Redirect("/")
	}

	logEntries := arn.FilterEditLogEntries(func(entry *arn.EditLogEntry) bool {
		return entry.UserID == user.ID
	})

	score := 0

	for _, entry := range logEntries {
		switch entry.Action {
		case "create":
			score += 10

		case "edit":
			score += 2

		case "delete", "arrayRemove":
			score++

		case "arrayAppend":
			// No score
		}
	}

	return ctx.HTML(components.Editor(ctx.URI(), score, user))
}
