package main

import (
	"io/ioutil"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/pages/airing"
	"github.com/animenotifier/notify.moe/pages/anime"
	"github.com/animenotifier/notify.moe/pages/dashboard"
	"github.com/animenotifier/notify.moe/pages/forum"
	"github.com/animenotifier/notify.moe/pages/forums"
	"github.com/animenotifier/notify.moe/pages/genre"
	"github.com/animenotifier/notify.moe/pages/genres"
	"github.com/animenotifier/notify.moe/pages/posts"
	"github.com/animenotifier/notify.moe/pages/profile"
	"github.com/animenotifier/notify.moe/pages/search"
	"github.com/animenotifier/notify.moe/pages/threads"
)

var app = aero.New()

func main() {
	// CSS
	app.SetStyle(components.CSS())

	// HTTPS
	app.Security.Certificate, _ = ioutil.ReadFile("security/fullchain.pem")
	app.Security.Key, _ = ioutil.ReadFile("security/privkey.pem")

	// Layout
	app.Layout = func(ctx *aero.Context, content string) string {
		return components.Layout(content)
	}

	// Ajax routes
	app.Ajax("/", dashboard.Get)
	app.Ajax("/anime", search.Get)
	app.Ajax("/anime/:id", anime.Get)
	app.Ajax("/genres", genres.Get)
	app.Ajax("/genres/:name", genre.Get)
	app.Ajax("/forum", forums.Get)
	app.Ajax("/forum/:tag", forum.Get)
	app.Ajax("/threads/:id", threads.Get)
	app.Ajax("/posts/:id", posts.Get)
	app.Ajax("/user/:nick", profile.Get)
	app.Ajax("/airing", airing.Get)

	app.Ajax("/test", func(ctx *aero.Context) string {
		return ctx.HTML(components.Test("Hello World"))
	})

	// Scripts
	scripts, _ := ioutil.ReadFile("temp/scripts.js")
	js := string(scripts)

	app.Get("/scripts.js", func(ctx *aero.Context) string {
		ctx.SetHeader("Content-Type", "application/javascript")
		return js
	})

	// For testing
	app.Get("/hello", func(ctx *aero.Context) string {
		return ctx.Text("Hello World")
	})

	// Let's go
	app.Run()
}
