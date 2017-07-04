package listimportanilist

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

func getMatches(ctx *aero.Context) ([]*arn.AniListMatch, string) {
	user := utils.GetUser(ctx)

	if user == nil {
		return nil, ctx.Error(http.StatusBadRequest, "Not logged in", nil)
	}

	authErr := arn.AniList.Authorize()

	if authErr != nil {
		return nil, ctx.Error(http.StatusBadRequest, "Couldn't authorize the Anime Notifier app on AniList", authErr)
	}

	allAnime, allErr := arn.AllAnime()

	if allErr != nil {
		return nil, ctx.Error(http.StatusBadRequest, "Couldn't load notify.moe list of all anime", allErr)
	}

	anilistAnimeList, err := arn.AniList.GetAnimeList(user)

	if err != nil {
		return nil, ctx.Error(http.StatusBadRequest, "Couldn't load your anime list from AniList", err)
	}

	matches := findAllMatches(allAnime, anilistAnimeList)

	return matches, ""
}

// Preview ...
func Preview(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	matches, response := getMatches(ctx)

	if response != "" {
		return response
	}

	return ctx.HTML(components.ImportAnilist(user, matches))
}

// Finish ...
func Finish(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	matches, response := getMatches(ctx)

	if response != "" {
		return response
	}

	animeList := user.AnimeList()

	for _, match := range matches {
		if match.ARNAnime == nil || match.AniListItem == nil {
			continue
		}

		item := &arn.AnimeListItem{
			AnimeID:  match.ARNAnime.ID,
			Status:   match.AniListItem.AnimeListStatus(),
			Episodes: match.AniListItem.EpisodesWatched,
			Notes:    match.AniListItem.Notes,
			Rating: &arn.AnimeRating{
				Overall: float64(match.AniListItem.ScoreRaw) / 10.0,
			},
			RewatchCount: match.AniListItem.Rewatched,
			Created:      arn.DateTimeUTC(),
			Edited:       arn.DateTimeUTC(),
		}

		animeList.Import(item)
	}

	err := animeList.Save()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error saving your anime list", err)
	}

	return ctx.Redirect("/+" + user.Nick + "/animelist")
}

// findAllMatches returns all matches for the anime inside an anilist anime list.
func findAllMatches(allAnime []*arn.Anime, animeList *arn.AniListAnimeList) []*arn.AniListMatch {
	matches := []*arn.AniListMatch{}

	matches = importList(matches, allAnime, animeList.Lists.Watching)
	matches = importList(matches, allAnime, animeList.Lists.Completed)
	matches = importList(matches, allAnime, animeList.Lists.PlanToWatch)
	matches = importList(matches, allAnime, animeList.Lists.OnHold)
	matches = importList(matches, allAnime, animeList.Lists.Dropped)

	custom, ok := animeList.CustomLists.(map[string][]*arn.AniListAnimeListItem)

	if !ok {
		return matches
	}

	for _, list := range custom {
		matches = importList(matches, allAnime, list)
	}

	return matches
}

// importList imports a single list inside an anilist anime list collection.
func importList(matches []*arn.AniListMatch, allAnime []*arn.Anime, animeListItems []*arn.AniListAnimeListItem) []*arn.AniListMatch {
	for _, item := range animeListItems {
		matches = append(matches, &arn.AniListMatch{
			AniListItem: item,
			ARNAnime:    arn.FindAniListAnime(item.Anime, allAnime),
		})
	}

	return matches
}
