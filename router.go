package main

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

func init() {
	app.Ajax("/", func(ctx *aero.Context) string {
		return ctx.HTML(components.Dashboard())
	})

	app.Ajax("/anime/:id", func(ctx *aero.Context) string {
		id, _ := ctx.GetInt("id")
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

	const threadsPerPage = 12

	forumHandler := func(ctx *aero.Context) string {
		tag := ctx.Get("tag")
		threads, _ := arn.GetThreadsByTag(tag)

		sort.Sort(threads)

		if len(threads) > threadsPerPage {
			threads = threads[:threadsPerPage]
		}

		for _, thread := range threads {
			thread.Author, _ = arn.GetUser(thread.AuthorID)
		}

		return ctx.HTML(components.Forum(threads))
	}

	app.Ajax("/forum", forumHandler)
	app.Ajax("/forum/:tag", forumHandler)

	app.Ajax("/threads/:id", func(ctx *aero.Context) string {
		id := ctx.Get("id")
		thread, err := arn.GetThread(id)

		if err != nil {
			return ctx.Text("Thread not found")
		}

		thread.Author, _ = arn.GetUser(thread.AuthorID)

		return ctx.HTML(components.Thread(thread))
	})
}
