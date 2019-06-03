package main

import (
	"fmt"
	"sort"

	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	list := []*arn.Anime{}

	for anime := range arn.StreamAnime() {
		if anime.Status != "current" || anime.Type != "tv" || anime.StartDate == "" || anime.StartDate < "2017-12" || anime.StartDate > "2018-02-01" {
			continue
		}

		list = append(list, anime)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Popularity.Total() > list[j].Popularity.Total()
	})

	for _, anime := range list {
		fmt.Printf("* [%s](/anime/%s)\n", anime.Title.Canonical, anime.ID)
	}
}
