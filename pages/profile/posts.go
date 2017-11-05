package profile

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const postLimit = 10

// GetPostsByUser shows all forum posts of a particular user.
func GetPostsByUser(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	posts := viewUser.Posts()
	arn.SortPostsLatestFirst(posts)

	var postables []arn.Postable

	if len(posts) >= postLimit {
		posts = posts[:postLimit]
	}

	postables = make([]arn.Postable, len(posts), len(posts))

	for i, post := range posts {
		postables[i] = arn.ToPostable(post)
	}

	return ctx.HTML(components.LatestPosts(postables, viewUser, utils.GetUser(ctx), ctx.URI()))

}
