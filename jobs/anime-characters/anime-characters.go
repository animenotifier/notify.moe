package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	anime, _ := arn.GetAnime("6887")
	err := anime.RefreshAnimeCharacters()
	arn.PanicOnError(err)

	color.Green("Finished.")
}
