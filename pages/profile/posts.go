package profile

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const postLimit = 10

// GetPostsbyUser shows all forum posts of a particular user.
func GetPostsByUser(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	user, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	posts := user.Posts()
	arn.SortPostsLatestLast(posts)

	var postables []arn.Postable

	if len(posts) >= postLimit {
		posts = posts[:postLimit]
	}

	postables = make([]arn.Postable, len(posts), len(posts))

	for i, post := range posts {

		postables[i] = arn.ToPostable(post)

	}

	return ctx.HTML(components.PostableList(postables))

}
