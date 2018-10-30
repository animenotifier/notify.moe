package main

import (
	"compress/gzip"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/animenotifier/mal"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/mal/parser"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Parsing MAL files")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	// Invoke via parameters
	if InvokeShellArgs() {
		return
	}

	filepath.Walk(path.Join(arn.Root, "jobs/mal-download/anime"), func(name string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(name, ".html.gz") {
			return nil
		}

		return readAnimeFile(name)
	})
}

func readAnimeFile(name string) error {
	file, err := os.Open(name)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer file.Close()

	reader, err := gzip.NewReader(file)

	if err != nil {
		fmt.Println(err)
		return err
	}

	anime, characters, err := malparser.ParseAnime(reader)

	if err != nil {
		fmt.Println(err)
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

	fmt.Println(anime.ID, anime.Title)
	arn.MAL.Set("Anime", anime.ID, anime)
	return nil
}

func readCharacterFile(name string) error {
	file, err := os.Open(name)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer file.Close()

	reader, err := gzip.NewReader(file)

	if err != nil {
		fmt.Println(err)
		return err
	}

	character, err := malparser.ParseCharacter(reader)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if character.ID == "" {
		return errors.New("Empty ID")
	}

	fmt.Println(character.ID, character.Name)
	arn.MAL.Set("Character", character.ID, character)
	return nil
}
