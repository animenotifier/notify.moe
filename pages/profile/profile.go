package profile

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxPosts = 5

// Get user profile page.
func Get(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(404, "User not found", err)
	}

	var user *arn.User
	var threads []*arn.Thread
	var animeList *arn.AnimeList

	aero.Parallel(func() {
		user = utils.GetUser(ctx)
	}, func() {
		animeList = viewUser.AnimeList()
	}, func() {
		threads = viewUser.Threads()

		arn.SortThreadsByDate(threads)

		if len(threads) > maxPosts {
			threads = threads[:maxPosts]
		}
	})

	return ctx.HTML(components.Profile(viewUser, user, animeList, threads))
}
