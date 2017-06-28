package main

import (
	"fmt"

	"github.com/animenotifier/arn"
)

var mappings = map[string]arn.Mapping{
	"13055": arn.Mapping{
		Service:   "shoboi/anime",
		ServiceID: "4528",
	},
}

func main() {
	for animeID, mapping := range mappings {
		anime, err := arn.GetAnime(animeID)

		if err != nil {
			panic(err)
		}

		fmt.Println(anime.ID, "=", mapping.Service, mapping.ServiceID)
		anime.AddMapping(mapping.Service, mapping.ServiceID)
		anime.Save()
	}
}
