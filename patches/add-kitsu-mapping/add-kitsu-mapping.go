package main

import (
	"fmt"
	"time"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Adding kitsu/anime mappings")

	defer color.Green("Finished")
	defer arn.Node.Close()

	for anime := range arn.StreamAnime() {
		fmt.Println(anime.ID, anime)
		anime.SetMapping("kitsu/anime", anime.ID)
		anime.Save()
	}

	time.Sleep(time.Second)
}
