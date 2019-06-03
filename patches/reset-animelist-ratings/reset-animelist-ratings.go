package main

import (
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	defer arn.Node.Close()

	animeList, err := arn.GetAnimeList("3eoFUr_zg")
	arn.PanicOnError(err)

	animeList.Lock()
	for _, item := range animeList.Items {
		item.Rating.Overall = 0
	}
	animeList.Unlock()

	animeList.Save()
}
