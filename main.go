package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"

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

	app.Get("/all", func(ctx *aero.Context) {
		var buffer bytes.Buffer

		results := make(chan *arn.Anime)
		arn.Scan("Anime", results)

		for anime := range results {
			buffer.WriteString(anime.Title.Romaji)
			buffer.WriteByte('\n')
		}

		ctx.Text(buffer.String())
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
