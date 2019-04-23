package main

import (
	"fmt"

	"github.com/animenotifier/arn"
	"github.com/akyoto/color"
)

func main() {
	color.Yellow("Checking anime status")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	for anime := range arn.StreamAnime() {
		if anime.Status != anime.CalculatedStatus() {
			fmt.Println("--------------------------------------------------------------------------------")
			fmt.Printf("%s (%s)\n", anime.Title.Canonical, anime.Type)
			fmt.Printf("%s => %s\n", color.RedString(anime.Status), color.YellowString(anime.CalculatedStatus()))
			fmt.Println(anime.ID)
		}
	}
}
