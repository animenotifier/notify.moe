package main

import (
	"io/ioutil"
	"strconv"

	"github.com/aerojs/aero"
	"github.com/blitzprog/arn"
	"github.com/valyala/fasthttp"
)

func main() {
	app := aero.New()

	cssBytes, _ := ioutil.ReadFile("layout.css")
	css := string(cssBytes)
	css = css

	app.Get("/", func(ctx *aero.Context) {
		ctx.Respond(Stream.Dashboard())
	})

	app.Get("/anime/:id", func(ctx *aero.Context) {
		id, _ := strconv.Atoi(ctx.Params.ByName("id"))
		anime, err := arn.GetAnime(id)
		anime = anime

		if err != nil {
			ctx.Respond("Anime not found")
			return
		}

		stream := fasthttp.AcquireByteBuffer()
		Stream.Layout(stream, anime, css)
		ctx.RespondBytes(stream.Bytes())
		Stream.Release(stream)
		// ctx.Respond(Render.Layout(Render.Anime(anime), css))
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
