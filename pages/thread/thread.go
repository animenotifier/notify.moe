package thread

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/server/middleware"
)

// Get thread.
func Get(ctx aero.Context) error {
	id := ctx.Get("id")
	user := arn.GetUserFromContext(ctx)

	// Fetch thread
	thread, err := arn.GetThread(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Thread not found", err)
	}

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = getOpenGraph(thread)
	return ctx.HTML(components.Thread(thread, user))
}
