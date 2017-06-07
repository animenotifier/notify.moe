package dashboard

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
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
	userID := ctx.Session().GetString("userId")

	if userID != "" {
		user, err := arn.GetUser(userID)

		if err != nil {
			return ctx.Error(http.StatusInternalServerError, "Error fetching user data", err)
		}

		return ctx.HTML("Welcome back, " + user.Nick + "!")
	}

	return ctx.HTML("ARN 4.0 is currently under construction.<br><a href='https://paypal.me/blitzprog' target='_blank' rel='noopener'>Support the development</a><br><a href='/auth/google'>Login via Google</a>")
}
