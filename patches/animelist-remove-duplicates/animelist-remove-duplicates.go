package main

import (
	"github.com/animenotifier/arn"
)

func main() {
	for animeList := range arn.StreamAnimeLists() {
		animeList.RemoveDuplicates()
		animeList.Save()
	}
}
