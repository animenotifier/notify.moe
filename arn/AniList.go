package arn

import (
	"fmt"

	"github.com/animenotifier/anilist"
)

// AniListAnimeFinder holds an internal map of ID to anime mappings
// and is therefore very efficient to use when trying to find
// anime by a given service and ID.
type AniListAnimeFinder struct {
	idToAnime    map[string]*Anime
	malIDToAnime map[string]*Anime
}

// NewAniListAnimeFinder creates a new finder for Anilist anime.
func NewAniListAnimeFinder() *AniListAnimeFinder {
	finder := &AniListAnimeFinder{
		idToAnime:    map[string]*Anime{},
		malIDToAnime: map[string]*Anime{},
	}

	for anime := range StreamAnime() {
		id := anime.GetMapping("anilist/anime")

		if id != "" {
			finder.idToAnime[id] = anime
		}

		malID := anime.GetMapping("myanimelist/anime")

		if malID != "" {
			finder.malIDToAnime[malID] = anime
		}
	}

	return finder
}

// GetAnime tries to find an AniList anime in our anime database.
func (finder *AniListAnimeFinder) GetAnime(id string, malID string) *Anime {
	animeByID, existsByID := finder.idToAnime[id]
	animeByMALID, existsByMALID := finder.malIDToAnime[malID]

	// Add anilist mapping to the MAL mapped anime if it's missing
	if existsByMALID && animeByMALID.GetMapping("anilist/anime") != id {
		animeByMALID.SetMapping("anilist/anime", id)
		animeByMALID.Save()

		finder.idToAnime[id] = animeByMALID
	}

	// If both MAL ID and AniList ID are matched, but the matched anime are different,
	// while the MAL IDs are different as well,
	// then we're trusting the MAL ID matching more and deleting the incorrect mapping.
	if existsByID && existsByMALID && animeByID.ID != animeByMALID.ID && animeByID.GetMapping("myanimelist/anime") != animeByMALID.GetMapping("myanimelist/anime") {
		animeByID.RemoveMapping("anilist/anime")
		animeByID.Save()

		delete(finder.idToAnime, id)

		fmt.Println("MAL / Anilist mismatch:")
		fmt.Println(animeByID.ID, animeByID)
		fmt.Println(animeByMALID.ID, animeByMALID)
	}

	if existsByID {
		return animeByID
	}

	if existsByMALID {
		return animeByMALID
	}

	return nil
}

// AniListAnimeListStatus returns the ARN version of the anime status.
func AniListAnimeListStatus(item *anilist.AnimeListItem) string {
	switch item.Status {
	case "CURRENT", "REPEATING":
		return AnimeListStatusWatching
	case "COMPLETED":
		return AnimeListStatusCompleted
	case "PLANNING":
		return AnimeListStatusPlanned
	case "PAUSED":
		return AnimeListStatusHold
	case "DROPPED":
		return AnimeListStatusDropped
	default:
		return AnimeListStatusPlanned
	}
}
