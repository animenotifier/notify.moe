package dashboard

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/flow"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/pages/frontpage"
	"github.com/animenotifier/notify.moe/utils"
)

const maxPosts = 5
const maxFollowing = 5

func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return frontpage.Get(ctx)
	}

	return Dashboard(ctx)
}

// Get dashboard.
func Dashboard(ctx *aero.Context) string {
	var user *arn.User
	var posts []*arn.Post
	var err error

	flow.Parallel(func() {
		user = utils.GetUser(ctx)
	}, func() {
		posts, err = arn.AllPostsSlice()
		arn.SortPostsLatestFirst(posts)

		if len(posts) > maxPosts {
			posts = posts[:maxPosts]
		}

	})

	followIDList := user.Following
	userList, err := arn.DB.GetMany("User", followIDList)

	if err != nil {
		return ctx.Error(500, "Error displaying dashboard", err)
	}

	followingList := userList.([]*arn.User)
	followingList = arn.SortByLastSeen(followingList)

	if len(followingList) > maxFollowing {
		followingList = followingList[:maxFollowing]
	}

	return ctx.HTML(components.Dashboard(posts, followingList))
}
