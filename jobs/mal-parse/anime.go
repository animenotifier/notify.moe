package main

import (
	"compress/gzip"
	"errors"
	"os"

	"github.com/akyoto/color"
	"github.com/animenotifier/mal"
	malparser "github.com/animenotifier/mal/parser"
	"github.com/animenotifier/notify.moe/arn"
)

// Read anime file
func readAnimeFile(name string) error {
	file, err := os.Open(name)

	if err != nil {
		color.Red(err.Error())
		return err
	}

	defer file.Close()

	reader, err := gzip.NewReader(file)

	if err != nil {
		color.Red(err.Error())
		return err
	}

	anime, characters, err := malparser.ParseAnime(reader)

	if err != nil {
		color.Red(err.Error())
		return err
	}

	if anime.ID == "" {
		return errors.New("Empty ID")
	}

	for _, character := range characters {
		obj, err := arn.MAL.Get("Character", character.ID)

		if err != nil {
			arn.MAL.Set("Character", character.ID, character)
			continue
		}

		existing := obj.(*mal.Character)
		modified := false

		if existing.Name != character.Name {
			existing.Name = character.Name
			modified = true
		}

		if existing.Image != character.Image {
			existing.Image = character.Image
			modified = true
		}

		if modified {
			arn.MAL.Set("Character", existing.ID, existing)
		}
	}

	// fmt.Println(anime.ID, anime.Title)
	arn.MAL.Set("Anime", anime.ID, anime)
	return nil
}
