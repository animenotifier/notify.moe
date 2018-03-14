package profile

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxPosts = 5
const maxTracks = 12

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
	user := utils.GetUser(ctx)
	animeList := viewUser.AnimeList()
	animeList.SortByRating()

	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":         viewUser.Nick,
			"og:image":         viewUser.AvatarLink("large"),
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

	return ctx.HTML(components.Profile(viewUser, user, animeList, ctx.URI()))
}
