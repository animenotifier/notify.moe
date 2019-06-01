package animelist

import (
	"net/http"
	"strconv"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/assets"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/middleware"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/animenotifier/notify.moe/utils/infinitescroll"
)

const (
	animeFirstLoad = 60
	animePerScroll = 40
)

// FilterByStatus returns a handler for the given anime list item status.
func FilterByStatus(status string) aero.Handler {
	return func(ctx aero.Context) error {
		user := utils.GetUser(ctx)
		return AnimeList(ctx, user, status)
	}
}

// AnimeList renders the anime list items.
func AnimeList(ctx aero.Context, user *arn.User, status string) error {
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

	statusList := animeList.FilterStatus(status)

	// Filter private items
	if user == nil || user.ID != viewUser.ID {
		statusList = statusList.WithoutPrivateItems()
	}

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
	return ctx.HTML(components.AnimeListPage(items, nextIndex, viewUser, user, status))
}
