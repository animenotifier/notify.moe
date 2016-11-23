package dashboard

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxPosts = 6

// Get ...
func Get(ctx *aero.Context) string {
	posts, err := arn.GetPosts()

	if err != nil {
		return ctx.Error(500, "Error fetching posts", err)
	}

	sort.Sort(sort.Reverse(posts))

	if len(posts) > maxPosts {
		posts = posts[:maxPosts]
	}

	return ctx.HTML(components.Dashboard(posts))
}
