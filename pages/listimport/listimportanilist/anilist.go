package listimportanilist

import (
	"net/http"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/anilist"
	"github.com/animenotifier/arn"
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

	return ctx.HTML(components.ImportAnilist(user, matches))
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
		if match.ARNAnime == nil || match.AniListItem == nil {
			continue
		}

		item := &arn.AnimeListItem{
			AnimeID:  match.ARNAnime.ID,
			Status:   arn.AniListAnimeListStatus(match.AniListItem),
			Episodes: match.AniListItem.EpisodesWatched,
			Notes:    match.AniListItem.Notes,
			Rating: arn.AnimeListItemRating{
				Overall: float64(match.AniListItem.ScoreRaw) / 10.0,
			},
			RewatchCount: match.AniListItem.Rewatched,
			Created:      arn.DateTimeUTC(),
			Edited:       arn.DateTimeUTC(),
		}

		animeList.Import(item)
	}

	animeList.Save()

	// Redirect
	prefix := "/"

	if strings.HasPrefix(ctx.URI(), "/_") {
		prefix = "/_/"
	}

	return ctx.Redirect(prefix + "+" + user.Nick + "/animelist/watching")
}

// getMatches finds and returns all matches for the logged in user.
func getMatches(ctx *aero.Context) ([]*arn.AniListMatch, string) {
	user := utils.GetUser(ctx)

	if user == nil {
		return nil, ctx.Error(http.StatusBadRequest, "Not logged in", nil)
	}

	authErr := anilist.Authorize()

	if authErr != nil {
		return nil, ctx.Error(http.StatusBadRequest, "Couldn't authorize the Anime Notifier app on AniList", authErr)
	}

	allAnime := arn.AllAnime()
	anilistAnimeList, err := anilist.GetAnimeList(user.Accounts.AniList.Nick)

	if err != nil {
		return nil, ctx.Error(http.StatusBadRequest, "Couldn't load your anime list from AniList", err)
	}

	matches := findAllMatches(allAnime, anilistAnimeList)

	return matches, ""
}

// findAllMatches returns all matches for the anime inside an anilist anime list.
func findAllMatches(allAnime []*arn.Anime, animeList *anilist.AnimeList) []*arn.AniListMatch {
	matches := []*arn.AniListMatch{}

	matches = importList(matches, allAnime, animeList.Lists.Watching)
	matches = importList(matches, allAnime, animeList.Lists.Completed)
	matches = importList(matches, allAnime, animeList.Lists.PlanToWatch)
	matches = importList(matches, allAnime, animeList.Lists.OnHold)
	matches = importList(matches, allAnime, animeList.Lists.Dropped)

	custom, ok := animeList.CustomLists.(map[string][]*anilist.AnimeListItem)

	if !ok {
		return matches
	}

	for _, list := range custom {
		matches = importList(matches, allAnime, list)
	}

	return matches
}

// importList imports a single list inside an anilist anime list collection.
func importList(matches []*arn.AniListMatch, allAnime []*arn.Anime, animeListItems []*anilist.AnimeListItem) []*arn.AniListMatch {
	for _, item := range animeListItems {
		matches = append(matches, &arn.AniListMatch{
			AniListItem: item,
			ARNAnime:    arn.FindAniListAnime(item.Anime, allAnime),
		})
	}

	return matches
}
