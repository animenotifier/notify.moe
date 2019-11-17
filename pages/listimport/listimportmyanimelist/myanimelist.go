package listimportmyanimelist

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aerogo/aero"
	"github.com/animenotifier/mal"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Preview shows an import preview.
func Preview(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in")
	}

	matches, err := getMatches(ctx)

	if err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}

	return ctx.HTML(components.ImportMyAnimeList(user, matches))
}

// Finish ...
func Finish(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in")
	}

	matches, err := getMatches(ctx)

	if err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}

	animeList := user.AnimeList()

	for _, match := range matches {
		if match.ARNAnime == nil || match.MyAnimeListItem == nil {
			continue
		}

		rewatchCount := 0

		if match.MyAnimeListItem.IsRewatching {
			rewatchCount = 1
		}

		item := &arn.AnimeListItem{
			AnimeID:  match.ARNAnime.ID,
			Status:   arn.MyAnimeListStatusToARNStatus(match.MyAnimeListItem.Status),
			Episodes: match.MyAnimeListItem.NumWatchedEpisodes,
			Notes:    "",
			Rating: arn.AnimeListItemRating{
				Overall: float64(match.MyAnimeListItem.Score),
			},
			RewatchCount: rewatchCount,
			Created:      arn.DateTimeUTC(),
			Edited:       arn.DateTimeUTC(),
		}

		animeList.Import(item)
	}

	animeList.Save()
	return ctx.HTML(components.ImportFinished(user))
}

// getMatches finds and returns all matches for the logged in user.
func getMatches(ctx aero.Context) ([]*arn.MyAnimeListMatch, error) {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return nil, errors.New("Not logged in")
	}

	malAnimeList, err := mal.GetAnimeList(user.Accounts.MyAnimeList.Nick)

	if err != nil {
		return nil, fmt.Errorf("Couldn't load your anime list from MyAnimeList: %s", err.Error())
	}

	matches := findAllMatches(malAnimeList)

	return matches, nil
}

// findAllMatches returns all matches for the anime inside an anilist anime list.
func findAllMatches(animeList mal.AnimeList) []*arn.MyAnimeListMatch {
	finder := arn.NewAnimeFinder("myanimelist/anime")
	matches := []*arn.MyAnimeListMatch{}

	for _, item := range animeList {
		matches = append(matches, &arn.MyAnimeListMatch{
			MyAnimeListItem: item,
			ARNAnime:        finder.GetAnime(strconv.Itoa(item.AnimeID)),
		})
	}

	return matches
}
