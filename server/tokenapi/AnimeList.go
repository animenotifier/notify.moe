package tokenapi

import (
	"errors"

	"github.com/animenotifier/notify.moe/arn"
)

// AnimeUpdate parses animeListEntry.out of the JSON and then tries to integrate them into the database
func AnimeUpdate(request *TokenRequest) error {
	animeListEntry := &arn.AnimeListItem{}
	animeJSON := request.JSON.Get("anime")

	animeListEntry.Status = animeJSON.Get("status").String()
	animeListEntry.AnimeID = animeJSON.Get("id").String()
	animeListEntry.Episodes = int(animeJSON.Get("episode").Int())
	animeListEntry.Notes = animeJSON.Get("notes").String()
	animeListEntry.RewatchCount = int(animeJSON.Get("rewatchCount").Int())
	animeListEntry.Private = animeJSON.Get("isPrivate").Bool()

	switch animeListEntry.Status {
	case arn.AnimeListStatusWatching:
	case arn.AnimeListStatusCompleted:
	case arn.AnimeListStatusPlanned:
	case arn.AnimeListStatusHold:
	case arn.AnimeListStatusDropped:
		break
	default:
		animeListEntry.Status = ""
	}

	if animeListEntry.AnimeID == "" {
		return errors.New("No anime ID has been supplied")
	}

	rating := animeJSON.Get("ratings")
	animeListEntry.Rating.Overall = rating.Get("overall").Float()
	animeListEntry.Rating.Story = rating.Get("story").Float()
	animeListEntry.Rating.Visuals = rating.Get("visuals").Float()
	animeListEntry.Rating.Soundtrack = rating.Get("soundtrack").Float()

	animeList := request.User.AnimeList()
	err := animeList.Import(animeListEntry)

	if err != nil {
		return err
	}

	return nil
}
