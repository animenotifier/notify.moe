package main

import (
	"fmt"
	"os"

	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	defer arn.Node.Close()

	oldID := os.Args[1]
	newID := os.Args[2]

	if oldID == "" || newID == "" {
		fmt.Println("Parameters: [old ID] [new ID]")
		os.Exit(1)
	}

	for animeList := range arn.StreamAnimeLists() {
		item := animeList.Find(oldID)

		if item == nil {
			continue
		}

		item.AnimeID = newID
		animeList.Save()
	}
}
