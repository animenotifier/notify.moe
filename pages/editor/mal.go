package editor

import (
	"sort"

	"github.com/animenotifier/notify.moe/utils/animediff"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/mal"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxCompareMALEntries = 20

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
		var differences []animediff.Difference

		// Canonical title
		if anime.Title.Canonical != malAnime.Title {
			hash := utils.HashString(malAnime.Title)

			if !arn.IsAnimeDifferenceIgnored(anime.ID, "mal", malAnime.ID, "CanonicalTitle", hash) {
				differences = append(differences, &animediff.CanonicalTitle{
					TitleA: anime.Title.Canonical,
					TitleB: malAnime.Title,
				})
			}
		}

		// Japanese title
		if anime.Title.Japanese != malAnime.JapaneseTitle {
			hash := utils.HashString(malAnime.JapaneseTitle)

			if !arn.IsAnimeDifferenceIgnored(anime.ID, "mal", malAnime.ID, "JapaneseTitle", hash) {
				differences = append(differences, &animediff.JapaneseTitle{
					TitleA: anime.Title.Japanese,
					TitleB: malAnime.JapaneseTitle,
				})
			}
		}

		// Synopsis
		if len(anime.Summary) < len(malAnime.Synopsis) {
			hash := utils.HashString(malAnime.Synopsis)

			if !arn.IsAnimeDifferenceIgnored(anime.ID, "mal", malAnime.ID, "Synopsis", hash) {
				differences = append(differences, &animediff.Synopsis{
					SynopsisA: anime.Summary,
					SynopsisB: malAnime.Synopsis,
				})
			}
		}

		// Compare genres
		hashA := utils.HashStringsNoOrder(anime.Genres)
		hashB := utils.HashStringsNoOrder(malAnime.Genres)

		if hashA != hashB {
			if !arn.IsAnimeDifferenceIgnored(anime.ID, "mal", malAnime.ID, "Genres", hashB) {
				differences = append(differences, &animediff.Genres{
					GenresA: anime.Genres,
					GenresB: malAnime.Genres,
				})
			}
		}

		// Add if there were any differences
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

	return ctx.HTML(components.CompareMAL(comparisons, ctx.URI(), user))
}
