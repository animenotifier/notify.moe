package dashboard

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/pages/frontpage"
	"github.com/animenotifier/notify.moe/utils"
)

const maxPosts = 5
const maxFollowing = 5

// Get dashboard.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return frontpage.Get(ctx)
	}

	posts, err := arn.GetPosts()

	if err != nil {
		return ctx.Error(500, "Error fetching posts", err)
	}

	arn.SortPostsLatestFirst(posts)

	if len(posts) > maxPosts {
		posts = posts[:maxPosts]
	}

	followIDList := user.Following
	var followingList []*arn.User

	if len(followIDList) > 0 {
		if len(followIDList) > maxFollowing {
			followIDList = followIDList[:maxFollowing]
		}

		userList, err := arn.DB.GetMany("User", followIDList)

		if err != nil {
			return ctx.Error(500, "Error fetching followed users", err)
		}

		followingList = userList.([]*arn.User)
	}

	return ctx.HTML(components.Dashboard(posts, followingList))
}
