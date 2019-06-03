package post

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/middleware"
	"github.com/animenotifier/notify.moe/utils"
)

// Get post.
func Get(ctx aero.Context) error {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)
	post, err := arn.GetPost(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Post not found", err)
	}

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = getOpenGraph(ctx, post)
	return ctx.HTML(components.Post(post, user))
}
