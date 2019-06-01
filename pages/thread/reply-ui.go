package thread

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// ReplyUI renders a new post area.
func ReplyUI(ctx aero.Context) error {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)
	thread, err := arn.GetThread(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Thread not found", err)
	}

	return ctx.HTML(components.NewPostArea(thread, user, "Reply") + components.NewPostActions(thread, true))
}
