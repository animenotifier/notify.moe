package post

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// ReplyUI renders a new post area.
func ReplyUI(ctx *aero.Context) string {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)
	post, err := arn.GetPost(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Post not found", err)
	}

	return ctx.HTML(components.NewPostArea(post, user, "Reply") + components.NewPostActions(post, true))
}
