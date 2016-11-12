package main

import (
	"strconv"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

func init() {
	app.Ajax("/", func(ctx *aero.Context) string {
		return ctx.HTML(components.Dashboard())
	})

	app.Ajax("/anime/:id", func(ctx *aero.Context) string {
		id, _ := strconv.Atoi(ctx.Get("id"))
		anime, err := arn.GetAnime(id)

		if err != nil {
			return ctx.Text("Anime not found")
		}

		return ctx.HTML(components.Anime(anime))
	})

	app.Ajax("/genres", func(ctx *aero.Context) string {
		return ctx.HTML(components.GenreOverview())
	})

	app.Ajax("/genres/:name", func(ctx *aero.Context) string {
		genreName := ctx.Get("name")
		genreInfo := new(arn.Genre)

		err := arn.GetObject("Genres", genreName, genreInfo)

		if err != nil {
			return err.Error()
		}

		return ctx.HTML(components.AnimeInGenre(genreInfo.Genre, genreInfo.AnimeList))
	})
}
