package listimportkitsu

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/kitsu"
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

	return ctx.HTML(components.ImportKitsu(user, matches))
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
		if match.ARNAnime == nil || match.KitsuItem == nil {
			continue
		}

		rating := match.KitsuItem.Attributes.RatingTwenty

		if rating < 0 {
			rating = 0
		}

		if rating > 20 {
			rating = 20
		}

		// Convert rating
		convertedRating := (float64(rating) / 20.0) * 10.0

		item := &arn.AnimeListItem{
			AnimeID:  match.ARNAnime.ID,
			Status:   arn.KitsuStatusToARNStatus(match.KitsuItem.Attributes.Status),
			Episodes: match.KitsuItem.Attributes.Progress,
			Notes:    match.KitsuItem.Attributes.Notes,
			Rating: arn.AnimeListItemRating{
				Overall: convertedRating,
			},
			RewatchCount: match.KitsuItem.Attributes.ReconsumeCount,
			Created:      arn.DateTimeUTC(),
			Edited:       arn.DateTimeUTC(),
		}

		animeList.Import(item)
	}

	animeList.Save()
	return ctx.HTML(components.ImportFinished(user))
}

// getMatches finds and returns all matches for the logged in user.
func getMatches(ctx aero.Context) ([]*arn.KitsuMatch, error) {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return nil, errors.New("Not logged in")
	}

	kitsuUser, err := kitsu.GetUser(user.Accounts.Kitsu.Nick)

	if err != nil {
		return nil, fmt.Errorf("Couldn't load your user info from Kitsu: %s", err.Error())
	}

	library := kitsuUser.StreamLibraryEntries()
	matches := findAllMatches(library)

	return matches, nil
}

// findAllMatches returns all matches for the anime inside an anilist anime list.
func findAllMatches(library chan *kitsu.LibraryEntry) []*arn.KitsuMatch {
	finder := arn.NewAnimeFinder("kitsu/anime")
	matches := []*arn.KitsuMatch{}

	for item := range library {
		// Ignore non-anime entries
		if item.Anime == nil {
			continue
		}

		matches = append(matches, &arn.KitsuMatch{
			KitsuItem: item,
			ARNAnime:  finder.GetAnime(item.Anime.ID),
		})
	}

	return matches
}
