package home

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/pages/frontpage"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/animenotifier/notify.moe/utils/infinitescroll"
)

const maxAnimeListItems = 50

// FilterByStatus returns a handler for the given anime list item status.
func FilterByStatus(status string) aero.Handle {
	return func(ctx *aero.Context) string {
		user := utils.GetUser(ctx)

		if user == nil {
			return frontpage.Get(ctx)
		}

		return AnimeListItems(ctx, user, status)
	}
}

// // AnimeList sends the anime list with the given status for given user.
// func AnimeList(ctx *aero.Context, user *arn.User, status string) string {
// 	viewUser := user
// 	animeList := viewUser.AnimeList()

// 	if animeList == nil {
// 		return ctx.Error(http.StatusNotFound, "Anime list not found", nil)
// 	}

// 	animeList = animeList.FilterStatus(status)
// 	animeList.Sort()
// 	items := animeList.Items

// 	if len(items) > maxAnimeListItems {
// 		items = items[:maxAnimeListItems]
// 	}

// 	fmt.Println(len(items))

// 	return ctx.HTML(components.Home(items, viewUser, user, status))
// }

// AnimeListItems renders the anime list items.
func AnimeListItems(ctx *aero.Context, user *arn.User, status string) string {
	viewUser := user
	index, _ := ctx.GetInt("index")

	// Fetch all eligible items
	animeList := viewUser.AnimeList()

	if animeList == nil {
		return ctx.Error(http.StatusNotFound, "Anime list not found", nil)
	}

	animeList = animeList.FilterStatus(status)

	// Sort the items
	animeList.Sort()

	// These are all animer list items for the given status
	allItems := animeList.Items

	// Slice the part that we need
	items := allItems[index:]

	if len(items) > maxAnimeListItems {
		items = items[:maxAnimeListItems]
	}

	// Next index
	nextIndex := infinitescroll.NextIndex(ctx, len(allItems), maxAnimeListItems, index)

	// In case we're scrolling, send items only (without the page frame)
	if index > 0 {
		return ctx.HTML(components.AnimeListScrollable(items, viewUser, user))
	}

	// Otherwise, send the full page
	return ctx.HTML(components.Home(items, nextIndex, viewUser, user, status))
}
