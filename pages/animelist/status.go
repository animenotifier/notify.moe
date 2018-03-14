package animelist

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

		return HomeAnimeList(ctx, user, status)
	}
}

// HomeAnimeList renders the anime list items.
func HomeAnimeList(ctx *aero.Context, user *arn.User, status string) string {
	nick := ctx.Get("nick")
	index, _ := ctx.GetInt("index")
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

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
	return ctx.HTML(components.HomeAnimeList(items, nextIndex, viewUser, user, status))
}
