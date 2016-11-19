package main

import (
	"io/ioutil"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/pages/anime"
	"github.com/animenotifier/notify.moe/pages/dashboard"
	"github.com/animenotifier/notify.moe/pages/forum"
	"github.com/animenotifier/notify.moe/pages/forums"
	"github.com/animenotifier/notify.moe/pages/genre"
	"github.com/animenotifier/notify.moe/pages/genres"
	"github.com/animenotifier/notify.moe/pages/threads"
)

var app = aero.New()

func main() {
	app.SetStyle(components.BundledCSS)

	scripts, _ := ioutil.ReadFile("temp/scripts.js")
	js := string(scripts)

	app.Get("/scripts.js", func(ctx *aero.Context) string {
		ctx.SetHeader("Content-Type", "application/javascript")
		return js
	})

	app.Get("/hello", func(ctx *aero.Context) string {
		return "Hello World"
	})

	app.Layout = func(ctx *aero.Context, content string) string {
		return components.Layout(content)
	}

	app.Ajax("/", dashboard.Get)
	app.Ajax("/anime/:id", anime.Get)
	app.Ajax("/genres", genres.Get)
	app.Ajax("/genres/:name", genre.Get)
	app.Ajax("/forum", forums.Get)
	app.Ajax("/forum/:tag", forum.Get)
	app.Ajax("/threads/:id", threads.Get)

	app.Run()
}
