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

const maxCompareMALEntries = 15

// CompareMAL ...
func CompareMAL(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	year := ctx.Get("year")
	status := ctx.Get("status")
	typ := ctx.Get("type")

	if year == "any" {
		year = ""
	}

	if status == "any" {
		status = ""
	}

	if typ == "any" {
		typ = ""
	}

	settings := user.Settings()
	settings.Editor.Filter.Year = year
	settings.Editor.Filter.Status = status
	settings.Editor.Filter.Type = typ
	settings.Save()

	animes := arn.FilterAnime(func(anime *arn.Anime) bool {
		if year != "" && (len(anime.StartDate) < 4 || anime.StartDate[:4] != year) {
			return false
		}

		if status != "" && anime.Status != status {
			return false
		}

		if typ != "" && anime.Type != typ {
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
					TitleA:      anime.Title.Canonical,
					TitleB:      malAnime.Title,
					NumericHash: hash,
				})
			}
		}

		// Japanese title
		if anime.Title.Japanese != malAnime.JapaneseTitle {
			hash := utils.HashString(malAnime.JapaneseTitle)

			if !arn.IsAnimeDifferenceIgnored(anime.ID, "mal", malAnime.ID, "JapaneseTitle", hash) {
				differences = append(differences, &animediff.JapaneseTitle{
					TitleA:      anime.Title.Japanese,
					TitleB:      malAnime.JapaneseTitle,
					NumericHash: hash,
				})
			}
		}

		// Romaji title
		if anime.Title.Romaji != malAnime.Title {
			hash := utils.HashString(malAnime.Title)

			if !arn.IsAnimeDifferenceIgnored(anime.ID, "mal", malAnime.ID, "RomajiTitle", hash) {
				differences = append(differences, &animediff.RomajiTitle{
					TitleA:      anime.Title.Romaji,
					TitleB:      malAnime.Title,
					NumericHash: hash,
				})
			}
		}

		// Airing start date
		if anime.StartDate != malAnime.StartDate && malAnime.StartDate != "" {
			hash := utils.HashString(malAnime.StartDate)

			if !arn.IsAnimeDifferenceIgnored(anime.ID, "mal", malAnime.ID, "StartDate", hash) {
				differences = append(differences, &animediff.StartDate{
					DateA:       anime.StartDate,
					DateB:       malAnime.StartDate,
					NumericHash: hash,
				})
			}
		}

		// Airing end date
		if anime.EndDate != malAnime.EndDate && malAnime.EndDate != "" {
			hash := utils.HashString(malAnime.EndDate)

			if !arn.IsAnimeDifferenceIgnored(anime.ID, "mal", malAnime.ID, "EndDate", hash) {
				differences = append(differences, &animediff.EndDate{
					DateA:       anime.EndDate,
					DateB:       malAnime.EndDate,
					NumericHash: hash,
				})
			}
		}

		// Status
		if anime.Status != malAnime.Status {
			hash := utils.HashString(malAnime.Status)

			if !arn.IsAnimeDifferenceIgnored(anime.ID, "mal", malAnime.ID, "Status", hash) {
				differences = append(differences, &animediff.Status{
					StatusA:     anime.Status,
					StatusB:     malAnime.Status,
					NumericHash: hash,
				})
			}
		}

		// EpisodeCount
		if anime.EpisodeCount != malAnime.EpisodeCount {
			hash := uint64(malAnime.EpisodeCount)

			if !arn.IsAnimeDifferenceIgnored(anime.ID, "mal", malAnime.ID, "EpisodeCount", hash) {
				differences = append(differences, &animediff.EpisodeCount{
					EpisodesA:   anime.EpisodeCount,
					EpisodesB:   malAnime.EpisodeCount,
					NumericHash: hash,
				})
			}
		}

		// Synopsis
		if len(anime.Summary) < 300 && len(anime.Summary)+50 < len(malAnime.Synopsis) {
			hash := utils.HashString(malAnime.Synopsis)

			if !arn.IsAnimeDifferenceIgnored(anime.ID, "mal", malAnime.ID, "Synopsis", hash) {
				differences = append(differences, &animediff.Synopsis{
					SynopsisA:   anime.Summary,
					SynopsisB:   malAnime.Synopsis,
					NumericHash: hash,
				})
			}
		}

		// Compare genres
		hashA := utils.HashStringsNoOrder(anime.Genres)
		hashB := utils.HashStringsNoOrder(malAnime.Genres)

		if hashA != hashB {
			if !arn.IsAnimeDifferenceIgnored(anime.ID, "mal", malAnime.ID, "Genres", hashB) {
				differences = append(differences, &animediff.Genres{
					GenresA:     anime.Genres,
					GenresB:     malAnime.Genres,
					NumericHash: hashB,
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

	return ctx.HTML(components.CompareMAL(comparisons, year, status, typ, "/editor/mal/diff/anime", user))
}
