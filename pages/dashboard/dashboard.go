package dashboard

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxPosts = 5

// Get ...
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user != nil {
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

	return ctx.HTML("ARN 4.0 is currently under construction.<br><a href='https://paypal.me/blitzprog' target='_blank' rel='noopener'>Support the development</a><br><a href='/auth/google'>Login via Google</a>")
}
