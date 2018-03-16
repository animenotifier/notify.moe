package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/mal/parser"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Importing MAL anime")

	defer arn.Node.Close()
	defer color.Green("Finished.")

	// readFile("../mal-download/files/anime-31240.html")

	filepath.Walk("../mal-download/files", func(name string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(name, ".html") {
			return nil
		}

		return readFile(name)
	})
}

func readFile(name string) error {
	file, err := os.Open(name)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	anime, err := malparser.ParseAnime(file)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if anime.ID == "" {
		return errors.New("Empty ID")
	}

	// fmt.Println(anime.ID, anime.Title)
	arn.MAL.Set("Anime", anime.ID, anime)
	return nil
}

// prettyPrint prints the object as indented JSON data on the console.
func prettyPrint(obj interface{}) {
	pretty, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Println(string(pretty))
}
