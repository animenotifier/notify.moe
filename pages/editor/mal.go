package editor

import (
	"sort"

	"github.com/OneOfOne/xxhash"
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/mal"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxCompareMALEntries = 10

// CompareMAL ...
func CompareMAL(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	year, _ := ctx.GetInt("year")
	animeType := ctx.Get("type")

	animes := arn.FilterAnime(func(anime *arn.Anime) bool {
		if year != 0 && year != anime.StartDateTime().Year() {
			return false
		}

		if animeType != "" && anime.Type != animeType {
			return false
		}

		return anime.GetMapping("myanimelist/anime") != ""
	})

	sort.Slice(animes, func(i, j int) bool {
		a := animes[i]
		b := animes[j]

		aPop := a.Popularity.Total()
		bPop := b.Popularity.Total()

		if aPop == bPop {
			return a.Title.Canonical < b.Title.Canonical
		}

		return aPop > bPop
	})

	comparisons := []*utils.MALComparison{}
	malAnimeCollection := arn.MAL.Collection("Anime")

	for _, anime := range animes {
		malID := anime.GetMapping("myanimelist/anime")
		obj, err := malAnimeCollection.Get(malID)

		if err != nil {
			continue
		}

		malAnime := obj.(*mal.Anime)
		var differences []utils.AnimeDiff

		sumA := uint64(0)

		for _, genre := range anime.Genres {
			h := xxhash.NewS64(0)
			h.Write([]byte(genre))
			numHash := h.Sum64()
			sumA += numHash
		}

		sumB := uint64(0)

		for _, genre := range malAnime.Genres {
			h := xxhash.NewS64(0)
			h.Write([]byte(genre))
			numHash := h.Sum64()
			sumB += numHash
		}

		if sumA != sumB {
			differences = append(differences, &utils.AnimeGenresDiff{
				GenresA: anime.Genres,
				GenresB: malAnime.Genres,
			})
		}

		if len(differences) > 0 {
			comparisons = append(comparisons, &utils.MALComparison{
				Anime:       anime,
				MALAnime:    malAnime,
				Differences: differences,
			})

			if len(comparisons) >= maxCompareMALEntries {
				break
			}
		}
	}

	return ctx.HTML(components.CompareMAL(comparisons, user))
}
