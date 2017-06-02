package dashboard

import (
	"github.com/aerogo/aero"
)

const maxPosts = 5

// Get ...
func Get(ctx *aero.Context) string {
	// posts, err := arn.GetPosts()

	// if err != nil {
	// 	return ctx.Error(500, "Error fetching posts", err)
	// }

	// arn.SortPostsLatestFirst(posts)

	// if len(posts) > maxPosts {
	// 	posts = posts[:maxPosts]
	// }

	// return ctx.HTML(components.Dashboard(posts))
	return ctx.HTML("ARN 4.0 is currently under construction.<br><a href='https://paypal.me/blitzprog' target='_blank' rel='noopener'>Support the development</a>")
}
