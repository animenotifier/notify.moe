package main

import (
	"compress/gzip"
	"errors"
	"os"

	"github.com/akyoto/color"
	malparser "github.com/animenotifier/mal/parser"
	"github.com/animenotifier/notify.moe/arn"
)

// Read character file
func readCharacterFile(name string) error {
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

	character, err := malparser.ParseCharacter(reader)

	if err != nil {
		color.Red(err.Error())
		return err
	}

	if character.ID == "" {
		return errors.New("Empty ID")
	}

	// fmt.Println(character.ID, character.Name)
	arn.MAL.Set("Character", character.ID, character)
	return nil
}
