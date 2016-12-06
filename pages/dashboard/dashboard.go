package dashboard

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxPosts = 5

// Get ...
func Get(ctx *aero.Context) string {
	posts, err := arn.GetPosts()

	if err != nil {
		return ctx.Error(500, "Error fetching posts", err)
	}

	arn.SortPostsLatestFirst(posts)

	if len(posts) > maxPosts {
		posts = posts[:maxPosts]
	}

	return ctx.HTML(components.Dashboard(posts))
}
