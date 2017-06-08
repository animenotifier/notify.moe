package main

import (
	"errors"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/utils"
)

func init() {
	// app.Get("/all/anime", func(ctx *aero.Context) string {
	// 	var titles []string

	// 	results := make(chan *arn.Anime)
	// 	arn.Scan("Anime", results)

	// 	for anime := range results {
	// 		titles = append(titles, anime.Title.Romaji)
	// 	}
	// 	sort.Strings(titles)

	// 	return ctx.Error(toString(len(titles)) + "\n\n" + strings.Join(titles, "\n"))
	// })

	app.Get("/api/anime/:id", func(ctx *aero.Context) string {
		id := ctx.Get("id")
		anime, err := arn.GetAnime(id)

		if err != nil {
			return ctx.Error(404, "Anime not found", err)
		}

		return ctx.JSON(anime)
	})

	app.Get("/api/users/:nick", func(ctx *aero.Context) string {
		nick := ctx.Get("nick")
		user, err := arn.GetUserByNick(nick)

		if err != nil {
			return ctx.Error(404, "User not found", err)
		}

		return ctx.JSON(user)
	})

	app.Get("/api/threads/:id", func(ctx *aero.Context) string {
		id := ctx.Get("id")
		thread, err := arn.GetThread(id)

		if err != nil {
			return ctx.Error(404, "Thread not found", err)
		}

		return ctx.JSON(thread)
	})

	app.Get("/api/anime/:id/add", func(ctx *aero.Context) string {
		animeID := ctx.Get("id")
		user := utils.GetUser(ctx)

		if user == nil {
			return ctx.Error(http.StatusBadRequest, "Not logged in", errors.New("User not logged in"))
		}

		animeList := user.AnimeList()

		if animeList.Contains(animeID) {
			return ctx.Error(http.StatusBadRequest, "Anime already added", errors.New("Anime has already been added"))
		}

		newItem := arn.AnimeListItem{
			AnimeID: animeID,
			Status:  arn.AnimeListStatusPlanned,
		}

		animeList.Items = append(animeList.Items, newItem)

		saveError := animeList.Save()

		if saveError != nil {
			return ctx.Error(http.StatusInternalServerError, "Could not save anime list in database", saveError)
		}

		return ctx.JSON(animeList)
	})
}
