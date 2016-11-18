package main

import (
	"sort"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

func init() {
	app.Get("/all/anime", func(ctx *aero.Context) string {
		var titles []string

		results := make(chan *arn.Anime)
		arn.Scan("Anime", results)

		for anime := range results {
			titles = append(titles, anime.Title.Romaji)
		}
		sort.Strings(titles)

		return ctx.Text(toString(len(titles)) + "\n\n" + strings.Join(titles, "\n"))
	})

	app.Get("/api/anime/:id", func(ctx *aero.Context) string {
		id, _ := ctx.GetInt("id")
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

	app.Get("/api/threads/:id", func(ctx *aero.Context) string {
		id := ctx.Get("id")
		thread, err := arn.GetThread(id)

		if err != nil {
			return ctx.Text("Thread not found")
		}

		return ctx.JSON(thread)
	})
}
