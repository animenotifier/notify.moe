package post

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get post.
func Get(ctx *aero.Context) string {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)
	post, err := arn.GetPost(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Post not found", err)
	}

	ctx.Data = getOpenGraph(ctx, post)
	return ctx.HTML(components.Post(post, user))
}
