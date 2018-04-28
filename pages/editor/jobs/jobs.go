package jobs

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

var running = map[string]bool{}
var lastRun = map[string]string{}

// Overview shows all background jobs.
func Overview(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	return ctx.HTML(components.EditorJobs(running, lastRun, ctx.URI(), user))
}
