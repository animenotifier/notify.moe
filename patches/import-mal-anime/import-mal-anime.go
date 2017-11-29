package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/animenotifier/mal"

	"github.com/aerogo/nano"
	"github.com/animenotifier/mal/parser"
)

var node = nano.New(5000)
var db = node.Namespace("mal").RegisterTypes((*mal.Anime)(nil))

func main() {
	defer node.Close()

	if true {
		filepath.Walk("data", func(name string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
				return err
			}

			if info.IsDir() {
				return nil
			}

			if !strings.HasSuffix(name, "-main.html") {
				return nil
			}

			return readFile(name)
		})
	} else {
		readFile("data/anime-1-main.html")
	}
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

	fmt.Println(anime.ID, anime.URL)
	db.Set("Anime", anime.ID, anime)
	return nil
}
