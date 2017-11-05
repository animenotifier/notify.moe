package profile

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/aerogo/flow"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxPosts = 5
const maxTracks = 3

// Get user profile page.
func Get(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(404, "User not found", err)
	}

	return Profile(ctx, viewUser)
}

// Profile renders the user profile page of the given viewUser.
func Profile(ctx *aero.Context, viewUser *arn.User) string {
	var user *arn.User
	var threads []*arn.Thread
	var animeList *arn.AnimeList
	var tracks []*arn.SoundTrack
	var posts []*arn.Post

	flow.Parallel(func() {
		user = utils.GetUser(ctx)
	}, func() {
		animeList = viewUser.AnimeList()

		// Sort by rating
		sort.Slice(animeList.Items, func(i, j int) bool {
			return animeList.Items[i].Rating.Overall > animeList.Items[j].Rating.Overall
		})
	})

	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":         viewUser.Nick,
			"og:image":         viewUser.LargeAvatar(),
			"og:url":           "https://" + ctx.App.Config.Domain + viewUser.Link(),
			"og:site_name":     "notify.moe",
			"og:description":   viewUser.Tagline,
			"og:type":          "profile",
			"profile:username": viewUser.Nick,
		},
		Meta: map[string]string{
			"description": viewUser.Tagline,
			"keywords":    viewUser.Nick + ",profile",
		},
	}

	ctx.Data = openGraph

	return ctx.HTML(components.Profile(viewUser, user, animeList, threads, posts, tracks, ctx.URI()))
}
