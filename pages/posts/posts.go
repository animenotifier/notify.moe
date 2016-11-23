package posts

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	id := ctx.Get("id")
	post, err := arn.GetPost(id)

	if err != nil {
		return ctx.Error(404, "Post not found", err)
	}

	return ctx.HTML(components.Post(post))
}
