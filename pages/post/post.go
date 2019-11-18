package post

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/server/middleware"
)

// Get post.
func Get(ctx aero.Context) error {
	id := ctx.Get("id")
	user := arn.GetUserFromContext(ctx)
	post, err := arn.GetPost(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Post not found", err)
	}

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = getOpenGraph(post)
	return ctx.HTML(components.Post(post, user))
}
