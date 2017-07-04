package listimportanilist

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get ...
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in", nil)
	}

	authErr := arn.AniList.Authorize()

	if authErr != nil {
		return ctx.Error(http.StatusBadRequest, "Couldn't authorize the Anime Notifier app on AniList", authErr)
	}

	allAnime, allErr := arn.AllAnime()

	if allErr != nil {
		return ctx.Error(http.StatusBadRequest, "Couldn't load notify.moe list of all anime", allErr)
	}

	animeList, err := arn.AniList.GetAnimeList(user)

	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Couldn't load your anime list from AniList", err)
	}

	matches := findAllMatches(allAnime, animeList)

	return ctx.HTML(components.ImportAnilist(user, matches))
}

// findAllMatches returns all matches for the anime inside an anilist anime list.
func findAllMatches(allAnime []*arn.Anime, animeList *arn.AniListAnimeList) []*arn.AniListMatch {
	matches := []*arn.AniListMatch{}

	matches = importList(matches, allAnime, animeList.Lists.Watching)
	matches = importList(matches, allAnime, animeList.Lists.Completed)
	matches = importList(matches, allAnime, animeList.Lists.PlanToWatch)
	matches = importList(matches, allAnime, animeList.Lists.OnHold)
	matches = importList(matches, allAnime, animeList.Lists.Dropped)

	for _, list := range animeList.CustomLists {
		matches = importList(matches, allAnime, list)
	}

	return matches
}

// importList imports a single list inside an anilist anime list collection.
func importList(matches []*arn.AniListMatch, allAnime []*arn.Anime, animeListItems []*arn.AniListAnimeListItem) []*arn.AniListMatch {
	for _, item := range animeListItems {
		matches = append(matches, &arn.AniListMatch{
			AniListAnime: item.Anime,
			ARNAnime:     arn.FindAniListAnime(item.Anime, allAnime),
		})
	}

	return matches
}
