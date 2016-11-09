package main

import (
	"strconv"

	"github.com/aerojs/aero"
	"github.com/animenotifier/arn"
)

func init() {
	app.Ajax("/", func(ctx *aero.Context) string {
		return ctx.HTML(Render.Dashboard())
	})

	app.Ajax("/anime/:id", func(ctx *aero.Context) string {
		id, _ := strconv.Atoi(ctx.Get("id"))
		anime, err := arn.GetAnime(id)

		if err != nil {
			return ctx.Text("Anime not found")
		}

		return ctx.HTML(Render.Anime(anime))
	})

	app.Ajax("/genres", func(ctx *aero.Context) string {
		return ctx.HTML(Render.GenreOverview())
	})

	app.Ajax("/genres/:name", func(ctx *aero.Context) string {
		genreName := ctx.Get("name")
		genreInfo := new(arn.Genre)

		err := arn.GetObject("Genres", genreName, genreInfo)

		if err != nil {
			return err.Error()
		}

		return ctx.HTML(Render.AnimeInGenre(genreInfo.Genre, genreInfo.AnimeList))
	})
}
