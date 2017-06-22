package profile

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// GetPostsbyUser shows all forum posts of a particular user.
func GetPostsByUser(ctx *aero.Context) string {
	postLimit := 10
	nick := ctx.Get("nick")
	user, err := arn.GetUserByNick(nick)
	var postables []arn.Postable
	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	posts := user.Posts()
	arn.SortPostsLatestLast(posts)

	if len(posts) >= postLimit {
		postables = make([]arn.Postable, postLimit, postLimit)
	} else {
		postables = make([]arn.Postable, len(posts), len(posts))
	}

	for i, post := range posts {

		if i == postLimit {
			break
		}

		postables[i] = arn.ToPostable(post)
	}

	return ctx.HTML(components.PostableList(postables))

}
