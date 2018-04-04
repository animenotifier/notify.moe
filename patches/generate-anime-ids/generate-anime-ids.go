package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Generating anime IDs")

	defer color.Green("Finished")
	defer arn.Node.Close()

	allAnime := arn.AllAnime()

	sort.Slice(allAnime, func(i, j int) bool {
		aID, _ := strconv.Atoi(allAnime[i].ID)
		bID, _ := strconv.Atoi(allAnime[j].ID)

		return aID < bID
	})

	for counter, anime := range allAnime {
		newID := arn.GenerateID("Anime")
		fmt.Printf("[%d / %d] Old [%s] New [%s] %s\n", counter+1, len(allAnime), color.YellowString(anime.ID), color.GreenString(newID), anime)
		anime.SetID(newID)
		time.Sleep(100 * time.Millisecond)
	}
}
