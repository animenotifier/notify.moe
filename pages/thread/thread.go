package thread

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get thread.
func Get(ctx *aero.Context) string {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)

	// Fetch thread
	thread, err := arn.GetThread(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Thread not found", err)
	}

	ctx.Data = getOpenGraph(ctx, thread)
	return ctx.HTML(components.Thread(thread, user))
}
