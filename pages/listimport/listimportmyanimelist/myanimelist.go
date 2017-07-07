package listimportmyanimelist

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/mal"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Preview shows an import preview.
func Preview(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in", nil)
	}

	matches, response := getMatches(ctx)

	if response != "" {
		return response
	}

	return ctx.HTML(components.ImportMyAnimeList(user, matches))
}

// Finish ...
func Finish(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in", nil)
	}

	return ctx.Redirect("/+" + user.Nick + "/animelist")
}

// getMatches finds and returns all matches for the logged in user.
func getMatches(ctx *aero.Context) ([]*arn.MyAnimeListMatch, string) {
	user := utils.GetUser(ctx)

	if user == nil {
		return nil, ctx.Error(http.StatusBadRequest, "Not logged in", nil)
	}

	malAnimeList, err := mal.GetAnimeList(user.Accounts.MyAnimeList.Nick)

	if err != nil {
		return nil, ctx.Error(http.StatusBadRequest, "Couldn't load your anime list from MyAnimeList", err)
	}

	matches := findAllMatches(malAnimeList)

	return matches, ""
}

// findAllMatches returns all matches for the anime inside an anilist anime list.
func findAllMatches(animeList *mal.AnimeList) []*arn.MyAnimeListMatch {
	matches := []*arn.MyAnimeListMatch{}

	for _, item := range animeList.Items {
		var anime *arn.Anime
		connection, err := arn.GetMyAnimeListToAnime(item.AnimeID)

		if err == nil {
			anime, _ = arn.GetAnime(connection.AnimeID)
		}

		matches = append(matches, &arn.MyAnimeListMatch{
			MyAnimeListItem: item,
			ARNAnime:        anime,
		})
	}

	return matches
}
