package main

import (
	"github.com/animenotifier/arn"
)

func main() {
	defer arn.Node.Close()

	for animeList := range arn.StreamAnimeLists() {
		animeList.RemoveDuplicates()
		animeList.Save()
	}
}
