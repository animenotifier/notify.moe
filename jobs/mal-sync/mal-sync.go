package main

import (
	"fmt"

	"github.com/aerogo/http/client"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/mal"
	"github.com/fatih/color"
)

const imageWidthThreshold = 225

var (
	malDB           = arn.Node.Namespace("mal").RegisterTypes((*mal.Anime)(nil))
	characterFinder = arn.NewCharacterFinder("myanimelist/character")
)

func main() {
	color.Yellow("Syncing with MAL DB")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	// Invoke via parameters
	if InvokeShellArgs() {
		return
	}

	// Sync the most important ones first
	allAnime := arn.AllAnime()
	arn.SortAnimeByQuality(allAnime)

	for _, anime := range allAnime {
		malID := anime.GetMapping("myanimelist/anime")

		if malID == "" {
			continue
		}

		sync(anime, malID)
	}
}

func sync(anime *arn.Anime, malID string) {
	obj, err := malDB.Get("Anime", malID)

	if err != nil {
		fmt.Println(err)
		return
	}

	malAnime := obj.(*mal.Anime)

	// Log title
	fmt.Printf("%s %s\n", color.CyanString(anime.Title.Canonical), malID)

	// Titles
	syncTitles(anime, malAnime)

	// Dates
	syncDates(anime, malAnime)

	// Episodes
	syncEpisodes(anime, malAnime)

	// Others
	syncOthers(anime, malAnime)

	// Characters
	syncCharacters(anime, malAnime)

	// Save in database
	anime.Save()
}

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
		anime.SetImageBytes(response.Bytes())
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

		if malCharacter.ID == "" || malCharacter.Name == "" || malCharacter.ImagePath == "" {
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
				character.DownloadImage(malCharacter.ImageLink())

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

func importCharacter(malCharacter *mal.Character) *arn.Character {
	fmt.Println("Importing MAL Character:", malCharacter.ID, malCharacter.Name, malCharacter.ImagePath)

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
