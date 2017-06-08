package profile

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxPosts = 5

// Get ...
func Get(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	viewUser, err := arn.GetUserByNick(nick)
	user := utils.GetUser(ctx)

	if err != nil {
		return ctx.Error(404, "User not found", err)
	}

	threads := viewUser.Threads()

	arn.SortThreadsByDate(threads)

	if len(threads) > maxPosts {
		threads = threads[:maxPosts]
	}

	animeList := viewUser.AnimeList()

	return ctx.HTML(components.Profile(viewUser, user, animeList, threads))
}
