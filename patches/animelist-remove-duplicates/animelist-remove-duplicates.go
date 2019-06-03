package main

import (
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	defer arn.Node.Close()

	for animeList := range arn.StreamAnimeLists() {
		animeList.RemoveDuplicates()
		animeList.Save()
	}
}
