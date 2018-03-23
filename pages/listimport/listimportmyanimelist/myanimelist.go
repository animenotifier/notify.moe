package listimportmyanimelist

import (
	"net/http"
	"strconv"

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

	matches, response := getMatches(ctx)

	if response != "" {
		return response
	}

	animeList := user.AnimeList()

	for _, match := range matches {
		if match.ARNAnime == nil || match.MyAnimeListItem == nil {
			continue
		}

		rating, _ := strconv.ParseFloat(match.MyAnimeListItem.MyScore, 64)
		episodesWatched, _ := strconv.Atoi(match.MyAnimeListItem.MyWatchedEpisodes)
		rewatchCount, convErr := strconv.Atoi(match.MyAnimeListItem.MyRewatching)

		if convErr != nil {
			rewatchCount = 0
		}

		item := &arn.AnimeListItem{
			AnimeID:  match.ARNAnime.ID,
			Status:   arn.MyAnimeListStatusToARNStatus(match.MyAnimeListItem.MyStatus),
			Episodes: episodesWatched,
			Notes:    "",
			Rating: &arn.AnimeRating{
				Overall: rating,
			},
			RewatchCount: rewatchCount,
			Created:      arn.DateTimeUTC(),
			Edited:       arn.DateTimeUTC(),
		}

		animeList.Import(item)
	}

	animeList.Save()

	return ctx.Redirect("/+" + user.Nick + "/animelist/watching")
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
	allAnime := arn.AllAnime()
	matches := []*arn.MyAnimeListMatch{}

	for _, item := range animeList.Items {
		matches = append(matches, &arn.MyAnimeListMatch{
			MyAnimeListItem: item,
			ARNAnime:        arn.FindMyAnimeListAnime(item.AnimeID, allAnime),
		})
	}

	return matches
}
