package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/aerojs/aero"
	"github.com/animenotifier/arn"
)

func main() {
	app := aero.New()

	cssBytes, _ := ioutil.ReadFile("layout.css")
	css := string(cssBytes)

	animeCSSBytes, _ := ioutil.ReadFile("anime.css")
	css += string(animeCSSBytes)

	app.SetStyle(css)

	scripts, _ := ioutil.ReadFile("scripts.js")
	js := string(scripts)

	// Define layout
	app.Layout = func(ctx *aero.Context, content string) string {
		return Render.Layout(content)
	}

	app.Register("/", func(ctx *aero.Context) string {
		return ctx.HTML(Render.Dashboard())
	})

	app.Register("/genres", func(ctx *aero.Context) string {
		return ctx.HTML(Render.GenreOverview())
	})

	type GenreInfo struct {
		Genre     string       `json:"genre"`
		AnimeList []*arn.Anime `json:"animeList"`
	}

	app.Register("/genres/:name", func(ctx *aero.Context) string {
		genreName := ctx.Params.ByName("name")
		genreInfo := new(GenreInfo)

		err := arn.GetObject("Genres", genreName, genreInfo)

		if err != nil {
			return err.Error()
		}

		return ctx.HTML(Render.AnimeInGenre(genreInfo.Genre, genreInfo.AnimeList))

		// var animeList []*arn.Anime
		// results := make(chan *arn.Anime)
		// arn.Scan("Anime", results)

		// for anime := range results {
		// 	genres := Map(anime.Genres, arn.FixGenre)
		// 	if Contains(genres, genreName) {
		// 		animeList = append(animeList, anime)
		// 	}
		// }

		// return ctx.HTML(Render.AnimeInGenre(genreName, animeList))
	})

	app.Register("/anime/:id", func(ctx *aero.Context) string {
		id, _ := strconv.Atoi(ctx.Params.ByName("id"))
		anime, err := arn.GetAnime(id)

		if err != nil {
			return ctx.Text("Anime not found")
		}

		return ctx.HTML(Render.Anime(anime))
	})

	// ---------------------------------------------------------------
	// API
	// ---------------------------------------------------------------

	app.Get("/scripts.js", func(ctx *aero.Context) string {
		ctx.SetHeader("Content-Type", "application/javascript")
		return js
	})

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
		id, _ := strconv.Atoi(ctx.Params.ByName("id"))
		anime, err := arn.GetAnime(id)

		if err != nil {
			return ctx.Text("Anime not found")
		}

		return ctx.JSON(anime)
	})

	app.Get("/api/users/:nick", func(ctx *aero.Context) string {
		nick := ctx.Params.ByName("nick")
		user, err := arn.GetUserByNick(nick)

		if err != nil {
			return ctx.Text("User not found")
		}

		return ctx.JSON(user)
	})

	app.Run()
}
