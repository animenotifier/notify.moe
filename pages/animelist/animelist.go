package animelist

import (
	"net/http"
	"strconv"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/pages/frontpage"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/animenotifier/notify.moe/utils/infinitescroll"
)

const (
	animeFirstLoad = 50
	animePerScroll = 15
)

// FilterByStatus returns a handler for the given anime list item status.
func FilterByStatus(status string) aero.Handle {
	return func(ctx *aero.Context) string {
		user := utils.GetUser(ctx)

		if user == nil {
			return frontpage.Get(ctx)
		}

		return AnimeList(ctx, user, status)
	}
}

// AnimeList renders the anime list items.
func AnimeList(ctx *aero.Context, user *arn.User, status string) string {
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

	statusList := animeList.FilterStatus(status)

	// Sort the items
	statusList.Sort()

	// These are all animer list items for the given status
	allItems := statusList.Items

	// Slice the part that we need
	items := allItems[index:]
	maxLength := animeFirstLoad

	if index > 0 {
		maxLength = animePerScroll
	}

	if len(items) > maxLength {
		items = items[:maxLength]
	}

	// Next index
	nextIndex := infinitescroll.NextIndex(ctx, len(allItems), maxLength, index)

	// OpenGraph data
	ctx.Data = &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       viewUser.Nick + "'s anime list",
			"og:image":       "https:" + viewUser.AvatarLink("large"),
			"og:url":         "https://" + ctx.App.Config.Domain + viewUser.Link(),
			"og:site_name":   "notify.moe",
			"og:description": strconv.Itoa(len(animeList.Items)) + " anime",

			// The OpenGraph type "profile" is meant for real-life persons but I think it's okay in this context.
			// An alternative would be to use "article" which is mostly used for blog posts and news.
			"og:type": "profile",
		},
		Meta: map[string]string{
			"description": viewUser.Nick + "'s anime list",
			"keywords":    "anime list",
		},
	}

	// In case we're scrolling, send items only (without the page frame)
	if index > 0 {
		return ctx.HTML(components.AnimeListScrollable(items, viewUser, user))
	}

	// Otherwise, send the full page
	return ctx.HTML(components.AnimeListPage(items, nextIndex, viewUser, user, status))
}
