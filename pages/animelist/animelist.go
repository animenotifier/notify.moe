package animelist

import (
	"net/http"
	"strconv"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/server/middleware"
	"github.com/animenotifier/notify.moe/utils/infinitescroll"
)

const (
	animeFirstLoad = 60
	animePerScroll = 40
)

// Filter filters a user's anime list item by the status.
func Filter(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	status := ctx.Get("status")
	sortBy := arn.SortByRating

	if user != nil {
		sortBy = user.Settings().SortBy
	}

	return AnimeList(ctx, user, status, sortBy)
}

// AnimeList renders the anime list items.
func AnimeList(ctx aero.Context, user *arn.User, status string, sortBy string) error {
	nick := ctx.Get("nick")
	index, _ := ctx.GetInt("index")
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	// Fetch all eligible items
	animeList := viewUser.AnimeList()

	if animeList == nil {
		return ctx.Error(http.StatusNotFound, "Anime list not found")
	}

	// Filter private items
	if user == nil || user.ID != viewUser.ID {
		animeList = animeList.WithoutPrivateItems()
	}

	statusLists := animeList.SplitByStatus()

	// Sort the items for the requested status only
	animeList = statusLists[status]
	animeList.Sort(sortBy)

	// These are all anime list items for the given status
	allItems := statusLists[status].Items

	// Slice the part that we need
	var items []*arn.AnimeListItem

	if index < len(allItems) {
		items = allItems[index:]
	}

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
	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       viewUser.Nick + "'s anime list",
			"og:image":       "https:" + viewUser.AvatarLink("large"),
			"og:url":         "https://" + assets.Domain + viewUser.Link(),
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
	return ctx.HTML(components.AnimeListPage(items, nextIndex, viewUser, user, statusLists))
}
