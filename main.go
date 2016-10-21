package main

import (
	"strconv"

	"github.com/aerojs/aero"
	"github.com/animenotifier/arn"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	arn.Init()

	app := aero.New()

	app.Get("/", func(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
		ctx.Write(ctx.RequestURI())
	})

	app.Get("/anime/:id", func(ctx *fasthttp.RequestCtx, params fasthttprouter.Params) {
		id, _ := strconv.Atoi(params.ByName("id"))
		anime, err := arn.GetAnime(id)

		if err != nil {
			ctx.WriteString("Anime not found")
			return
		}

		ctx.WriteString(anime.Description)
	})

	app.Run()
}
