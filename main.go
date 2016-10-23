package main

import (
	"io/ioutil"
	"strconv"

	"github.com/aerojs/aero"
	"github.com/blitzprog/arn"
)

func main() {
	app := aero.New()

	example, _ := ioutil.ReadFile("security/frontpage.html")

	app.Get("/", func(ctx *aero.Context) {
		ctx.RespondBytes(example)
	})

	app.Get("/hello", func(ctx *aero.Context) {
		ctx.Respond("Hello World")
	})

	template := aero.NewTemplate("pages/anime/anime.pug")

	app.Get("/anime/:id", func(ctx *aero.Context) {
		id, _ := strconv.Atoi(ctx.Params.ByName("id"))
		anime, err := arn.GetAnime(id)

		if err != nil {
			ctx.Respond("Anime not found")
			return
		}

		templateParams := make(map[string]interface{})
		templateParams["anime"] = anime
		ctx.Respond(template.Render(templateParams))
	})

	app.Run()
}
