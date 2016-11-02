package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/aerojs/aero"
	"github.com/animenotifier/arn"
)

func s(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func main() {
	app := aero.New()

	cssBytes, _ := ioutil.ReadFile("layout.css")
	css := string(cssBytes)

	animeCSSBytes, _ := ioutil.ReadFile("anime.css")
	css += string(animeCSSBytes)

	scripts, _ := ioutil.ReadFile("scripts.js")
	js := string(scripts)

	app.Get("/", func(ctx *aero.Context) {
		ctx.HTML(Render.Layout(Render.Dashboard(), css))
	})

	app.Get("/_/", func(ctx *aero.Context) {
		ctx.HTML(Render.Dashboard())
	})

	app.Get("/all/anime", func(ctx *aero.Context) {
		start := time.Now()
		var titles []string

		results := make(chan *arn.Anime)
		arn.Scan("Anime", results)

		for anime := range results {
			titles = append(titles, anime.Title.Romaji)
		}
		sort.Strings(titles)

		ctx.Text(s(len(titles)) + " anime fetched in " + s(time.Since(start)) + "\n\n" + strings.Join(titles, "\n"))
	})

	app.Get("/anime/:id", func(ctx *aero.Context) {
		id, _ := strconv.Atoi(ctx.Params.ByName("id"))
		anime, err := arn.GetAnime(id)

		if err != nil {
			ctx.Text("Anime not found")
			return
		}

		ctx.HTML(Render.Layout(Render.Anime(anime), css))
	})

	app.Get("/_/anime/:id", func(ctx *aero.Context) {
		id, _ := strconv.Atoi(ctx.Params.ByName("id"))
		anime, err := arn.GetAnime(id)

		if err != nil {
			ctx.Text("Anime not found")
			return
		}

		ctx.HTML(Render.Anime(anime))
	})

	app.Get("/api/anime/:id", func(ctx *aero.Context) {
		id, _ := strconv.Atoi(ctx.Params.ByName("id"))
		anime, err := arn.GetAnime(id)

		if err != nil {
			ctx.Text("Anime not found")
			return
		}

		ctx.JSON(anime)
	})

	app.Get("/api/users/:nick", func(ctx *aero.Context) {
		nick := ctx.Params.ByName("nick")
		user, err := arn.GetUserByNick(nick)

		if err != nil {
			ctx.Text("User not found")
			return
		}

		ctx.JSON(user)
	})

	app.Get("/genres", func(ctx *aero.Context) {
		ctx.HTML(Render.Layout(Render.GenreOverview(), css))
	})

	app.Get("/_/genres", func(ctx *aero.Context) {
		ctx.HTML(Render.GenreOverview())
	})

	app.Get("/scripts.js", func(ctx *aero.Context) {
		ctx.SetHeader("Content-Type", "application/javascript")
		ctx.Respond(js)
	})

	// layout := aero.NewTemplate("layout.pug")
	// template := aero.NewTemplate("anime.pug")

	// app.Get("/anime/:id", func(ctx *aero.Context) {
	// 	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	// 	anime, err := arn.GetAnime(id)

	// 	if err != nil {
	// 		ctx.Respond("Anime not found")
	// 		return
	// 	}

	// 	content := template.Render(map[string]interface{}{
	// 		"anime": anime,
	// 	})

	// 	final := layout.Render(map[string]interface{}{
	// 		"content": content,
	// 	})

	// 	final = strings.Replace(final, cssSearch, cssReplace, 1)

	// 	ctx.Respond(final)
	// })

	// app.Get("/t", func(ctx *aero.Context) {
	// 	ctx.Respond(templates.Hello("abc"))
	// })

	app.Run()
}
