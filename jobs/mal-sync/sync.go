package main

import (
	"fmt"

	"github.com/aerogo/http/client"
	"github.com/akyoto/color"
	"github.com/animenotifier/mal"
	"github.com/animenotifier/notify.moe/arn"
)

// Sync titles
func syncTitles(anime *arn.Anime, malAnime *mal.Anime) {
	if anime.Title.Japanese == "" && malAnime.JapaneseTitle != "" {
		fmt.Println("JapaneseTitle:", malAnime.JapaneseTitle)
		anime.Title.Japanese = malAnime.JapaneseTitle
	}

	if anime.Title.English == "" && malAnime.EnglishTitle != "" {
		fmt.Println("EnglishTitle:", malAnime.EnglishTitle)
		anime.Title.English = malAnime.EnglishTitle
	}
}

// Sync dates
func syncDates(anime *arn.Anime, malAnime *mal.Anime) {
	if anime.StartDate == "" && malAnime.StartDate != "" {
		fmt.Println("StartDate:", malAnime.StartDate)
		anime.StartDate = malAnime.StartDate
	}

	if anime.EndDate == "" && malAnime.EndDate != "" {
		fmt.Println("EndDate:", malAnime.EndDate)
		anime.EndDate = malAnime.EndDate
	}
}

// Sync episodes
func syncEpisodes(anime *arn.Anime, malAnime *mal.Anime) {
	if anime.EpisodeCount == 0 && malAnime.EpisodeCount != 0 {
		fmt.Println("EpisodeCount:", malAnime.EpisodeCount)
		anime.EpisodeCount = malAnime.EpisodeCount
	}

	if anime.EpisodeLength == 0 && malAnime.EpisodeLength != 0 {
		fmt.Println("EpisodeLength:", malAnime.EpisodeLength)
		anime.EpisodeLength = malAnime.EpisodeLength
	}
}

// Sync others
func syncOthers(anime *arn.Anime, malAnime *mal.Anime) {
	if len(anime.Genres) == 0 && len(malAnime.Genres) > 0 {
		fmt.Println("Genres:", malAnime.Genres)
		anime.Genres = malAnime.Genres
	}

	if anime.Source == "" && malAnime.Source != "" {
		fmt.Println("Source:", malAnime.Source)
		anime.Source = malAnime.Source
	}
}

// Sync image
func syncImage(anime *arn.Anime, malAnime *mal.Anime) {
	if (anime.HasImage() && anime.Image.Width > imageWidthThreshold) || malAnime.Image == "" {
		return
	}

	fmt.Println("Downloading image:", malAnime.Image)
	response, err := client.Get(malAnime.Image).End()

	if err == nil && response.StatusCode() == 200 {
		err := anime.SetImageBytes(response.Bytes())

		if err != nil {
			color.Red("Error setting image: %s", err.Error())
		}
	} else {
		color.Red("Error downloading image")
	}
}

// Sync characters
func syncCharacters(anime *arn.Anime, malAnime *mal.Anime) {
	// Check for existence of characters
	animeCharacters := anime.Characters()
	modifiedCharacters := false

	for _, malAnimeCharacter := range malAnime.Characters {
		// Make sure we have no invalid entries
		if malAnimeCharacter.ID == "" || malAnimeCharacter.Role == "" {
			fmt.Println("Skip:", malAnimeCharacter)
			continue
		}

		animeCharacter := animeCharacters.FindByMapping("myanimelist/character", malAnimeCharacter.ID)

		if animeCharacter != nil {
			continue
		}

		obj, err := malDB.Get("Character", malAnimeCharacter.ID)

		// If we don't have the MAL character in the DB,
		// we can't import anything here.
		if err != nil {
			continue
		}

		malCharacter := obj.(*mal.Character)

		if malCharacter.ID == "" || malCharacter.Name == "" || malCharacter.Image == "" {
			fmt.Println("Skip character:", malAnimeCharacter.ID)
			continue
		}

		fmt.Println("Importing MAL AnimeCharacter:", malAnimeCharacter.ID, "as", malAnimeCharacter.Role)

		// Import character if needed
		character := characterFinder.GetCharacter(malAnimeCharacter.ID)

		if character == nil {
			character = importCharacter(malCharacter)
		} else {
			fmt.Println("Found existing character:", character)

			// Download image if missing
			if !character.HasImage() {
				fmt.Println("Downloading missing image for character:", character)
				err := character.DownloadImage(malCharacter.ImageLink())

				// Cancel import if that character has no image
				if err != nil {
					color.Red(err.Error())
					continue
				}
			}
		}

		// If import failed, continue
		if character == nil {
			continue
		}

		// Add to anime characters
		err = animeCharacters.Add(&arn.AnimeCharacter{
			CharacterID: character.ID,
			Role:        malAnimeCharacter.Role,
		})

		if err != nil {
			color.Red(err.Error())
		}

		modifiedCharacters = true
	}

	if modifiedCharacters {
		animeCharacters.Save()
	}
}

// import character
func importCharacter(malCharacter *mal.Character) *arn.Character {
	fmt.Println("Importing MAL Character:", malCharacter.ID, malCharacter.Name, malCharacter.Image)

	character := arn.NewCharacter()
	character.Name.Canonical = malCharacter.Name

	// Cancel the import if image could not be fetched
	err := character.DownloadImage(malCharacter.ImageLink())

	if err != nil {
		color.Red(err.Error())
		return nil
	}

	// Add mapping
	character.SetMapping("myanimelist/character", malCharacter.ID)

	// Save character in DB
	character.Save()

	// Add to character finder so we don't create duplicates of this character
	characterFinder.Add(character)

	return character
}
