package editor

import (
	"github.com/akyoto/hash"
	"github.com/animenotifier/mal"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/animenotifier/notify.moe/utils/animediff"
)

// diff titles
func diffTitles(anime *arn.Anime, malAnime *mal.Anime) []animediff.Difference {
	var differences []animediff.Difference

	// Canonical title
	if anime.Title.Canonical != malAnime.Title {
		hash := hash.String(malAnime.Title)

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
		hash := hash.String(malAnime.JapaneseTitle)

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
		hash := hash.String(malAnime.Title)

		if !arn.IsAnimeDifferenceIgnored(anime.ID, "mal", malAnime.ID, "RomajiTitle", hash) {
			differences = append(differences, &animediff.RomajiTitle{
				TitleA:      anime.Title.Romaji,
				TitleB:      malAnime.Title,
				NumericHash: hash,
			})
		}
	}

	return differences
}

// diff dates
func diffDates(anime *arn.Anime, malAnime *mal.Anime) []animediff.Difference {
	var differences []animediff.Difference

	// Airing start date
	if anime.StartDate != malAnime.StartDate && malAnime.StartDate != "" {
		hash := hash.String(malAnime.StartDate)

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
		hash := hash.String(malAnime.EndDate)

		if !arn.IsAnimeDifferenceIgnored(anime.ID, "mal", malAnime.ID, "EndDate", hash) {
			differences = append(differences, &animediff.EndDate{
				DateA:       anime.EndDate,
				DateB:       malAnime.EndDate,
				NumericHash: hash,
			})
		}
	}

	return differences
}

// diff episodes
func diffEpisodes(anime *arn.Anime, malAnime *mal.Anime) []animediff.Difference {
	var differences []animediff.Difference

	// EpisodeCount
	if malAnime.EpisodeCount != 0 && anime.EpisodeCount != malAnime.EpisodeCount {
		hash := uint64(malAnime.EpisodeCount)

		if !arn.IsAnimeDifferenceIgnored(anime.ID, "mal", malAnime.ID, "EpisodeCount", hash) {
			differences = append(differences, &animediff.EpisodeCount{
				EpisodesA:   anime.EpisodeCount,
				EpisodesB:   malAnime.EpisodeCount,
				NumericHash: hash,
			})
		}
	}

	return differences
}

// diff status
func diffStatus(anime *arn.Anime, malAnime *mal.Anime) []animediff.Difference {
	var differences []animediff.Difference

	// Status
	if anime.Status != malAnime.Status {
		hash := hash.String(malAnime.Status)

		if !arn.IsAnimeDifferenceIgnored(anime.ID, "mal", malAnime.ID, "Status", hash) {
			differences = append(differences, &animediff.Status{
				StatusA:     anime.Status,
				StatusB:     malAnime.Status,
				NumericHash: hash,
			})
		}
	}

	return differences
}

// diff synopsis
func diffSynopsis(anime *arn.Anime, malAnime *mal.Anime) []animediff.Difference {
	var differences []animediff.Difference

	// Synopsis
	if len(anime.Summary) < 300 && len(anime.Summary)+50 < len(malAnime.Synopsis) {
		hash := hash.String(malAnime.Synopsis)

		if !arn.IsAnimeDifferenceIgnored(anime.ID, "mal", malAnime.ID, "Synopsis", hash) {
			differences = append(differences, &animediff.Synopsis{
				SynopsisA:   anime.Summary,
				SynopsisB:   malAnime.Synopsis,
				NumericHash: hash,
			})
		}
	}

	return differences
}

// diff genres
func diffGenres(anime *arn.Anime, malAnime *mal.Anime) []animediff.Difference {
	var differences []animediff.Difference

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

	return differences
}
