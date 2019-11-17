package listimportanilist

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aerogo/aero"
	"github.com/animenotifier/anilist"
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

	return ctx.HTML(components.ImportAnilist(user, matches))
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
		if match.ARNAnime == nil || match.AniListItem == nil {
			continue
		}

		item := &arn.AnimeListItem{
			AnimeID:  match.ARNAnime.ID,
			Status:   arn.AniListAnimeListStatus(match.AniListItem),
			Episodes: match.AniListItem.Progress,
			Notes:    match.AniListItem.Notes,
			Rating: arn.AnimeListItemRating{
				Overall: float64(match.AniListItem.ScoreRaw) / 10.0,
			},
			RewatchCount: match.AniListItem.Repeat,
			Private:      match.AniListItem.Private,
			Created:      arn.DateTimeUTC(),
			Edited:       arn.DateTimeUTC(),
		}

		animeList.Import(item)
	}

	animeList.Save()
	return ctx.HTML(components.ImportFinished(user))
}

// getMatches finds and returns all matches for the logged in user.
func getMatches(ctx aero.Context) ([]*arn.AniListMatch, error) {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return nil, errors.New("Not logged in")
	}

	anilistUser, err := anilist.GetUser(user.Accounts.AniList.Nick)

	if err != nil {
		return nil, fmt.Errorf("User doesn't exist on AniList: %s", err.Error())
	}

	anilistAnimeList, err := anilist.GetAnimeList(anilistUser.ID)

	if err != nil {
		return nil, fmt.Errorf("Couldn't load your anime list from AniList: %s", err.Error())
	}

	// Find matches
	matches := findAllMatches(anilistAnimeList)

	return matches, nil
}

// findAllMatches returns all matches for the anime inside an anilist anime list.
func findAllMatches(animeList *anilist.AnimeList) []*arn.AniListMatch {
	finder := arn.NewAniListAnimeFinder()
	matches := []*arn.AniListMatch{}

	for _, list := range animeList.Lists {
		matches = importList(matches, finder, list.Entries)
	}

	return matches
}

// importList imports a single list inside an anilist anime list collection.
func importList(matches []*arn.AniListMatch, finder *arn.AniListAnimeFinder, animeListItems []*anilist.AnimeListItem) []*arn.AniListMatch {
	for _, item := range animeListItems {
		matches = append(matches, &arn.AniListMatch{
			AniListItem: item,
			ARNAnime:    finder.GetAnime(strconv.Itoa(item.Anime.ID), strconv.Itoa(item.Anime.MALID)),
		})
	}

	return matches
}
