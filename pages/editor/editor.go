package editor

import (
	"net/http"
	"strconv"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get ...
func Get(ctx aero.Context) error {
	user := utils.GetUser(ctx)

	if user == nil || (user.Role != "admin" && user.Role != "editor") {
		return ctx.Redirect(http.StatusFound, "/")
	}

	ignoreDifferences := arn.FilterIgnoreAnimeDifferences(func(entry *arn.IgnoreAnimeDifference) bool {
		return entry.CreatedBy == user.ID
	})

	score := len(ignoreDifferences) * arn.IgnoreAnimeDifferenceEditorScore
	scoreTypes := map[string]int{}

	logEntries := arn.FilterEditLogEntries(func(entry *arn.EditLogEntry) bool {
		return entry.UserID == user.ID
	})

	for _, entry := range logEntries {
		entryScore := entry.EditorScore()
		score += entryScore

		if entry.ObjectType != "" {
			scoreTypes[entry.ObjectType]++
		}
	}

	scoreTitle := ""

	for objectType, score := range scoreTypes {
		scoreTitle += objectType + ": " + strconv.Itoa(score) + "\n"
	}

	return ctx.HTML(components.Editor(ctx.Path(), score, scoreTitle, scoreTypes, user))
}
