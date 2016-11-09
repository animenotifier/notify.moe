package main

import (
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/aerojs/aero"
	"github.com/animenotifier/arn"
)

func init() {
	app.Get("/all/anime", func(ctx *aero.Context) string {
		start := time.Now()
		var titles []string

		results := make(chan *arn.Anime)
		arn.Scan("Anime", results)

		for anime := range results {
			titles = append(titles, anime.Title.Romaji)
		}
		sort.Strings(titles)

		return ctx.Text(s(len(titles)) + " anime fetched in " + s(time.Since(start)) + "\n\n" + strings.Join(titles, "\n"))
	})

	app.Get("/api/anime/:id", func(ctx *aero.Context) string {
		id, _ := strconv.Atoi(ctx.Get("id"))
		anime, err := arn.GetAnime(id)

		if err != nil {
			return ctx.Text("Anime not found")
		}

		return ctx.JSON(anime)
	})

	app.Get("/api/users/:nick", func(ctx *aero.Context) string {
		nick := ctx.Get("nick")
		user, err := arn.GetUserByNick(nick)

		if err != nil {
			return ctx.Text("User not found")
		}

		return ctx.JSON(user)
	})
}
