package main

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/pages/airing"
	"github.com/animenotifier/notify.moe/pages/anime"
	"github.com/animenotifier/notify.moe/pages/awards"
	"github.com/animenotifier/notify.moe/pages/dashboard"
	"github.com/animenotifier/notify.moe/pages/forum"
	"github.com/animenotifier/notify.moe/pages/forums"
	"github.com/animenotifier/notify.moe/pages/posts"
	"github.com/animenotifier/notify.moe/pages/profile"
	"github.com/animenotifier/notify.moe/pages/search"
	"github.com/animenotifier/notify.moe/pages/threads"
	"github.com/animenotifier/notify.moe/pages/users"
)

var app = aero.New()

func main() {
	// HTTPS
	app.Security.Load("security/fullchain.pem", "security/privkey.pem")

	// CSS
	app.SetStyle(components.CSS())

	// Session store
	app.Sessions.Store = arn.NewAerospikeStore("Session")

	// Layout
	app.Layout = func(ctx *aero.Context, content string) string {
		return components.Layout(content)
	}

	// Ajax routes
	app.Ajax("/", dashboard.Get)
	app.Ajax("/anime", search.Get)
	app.Ajax("/anime/:id", anime.Get)
	app.Ajax("/forum", forums.Get)
	app.Ajax("/forum/:tag", forum.Get)
	app.Ajax("/threads/:id", threads.Get)
	app.Ajax("/posts/:id", posts.Get)
	app.Ajax("/user/:nick", profile.Get)
	app.Ajax("/user/:nick/threads", threads.GetByUser)
	app.Ajax("/users", users.Get)
	app.Ajax("/airing", airing.Get)
	app.Ajax("/awards", awards.Get)
	// app.Ajax("/genres", genres.Get)
	// app.Ajax("/genres/:name", genre.Get)

	// Favicon
	app.Get("/favicon.ico", func(ctx *aero.Context) string {
		return ctx.File("images/icons/favicon.ico")
	})

	// Scripts
	app.Get("/scripts.js", func(ctx *aero.Context) string {
		return ctx.File("temp/scripts.js")
	})

	// Web manifest
	app.Get("/manifest.json", func(ctx *aero.Context) string {
		return ctx.JSON(app.Config.Manifest)
	})

	// Cover image
	app.Get("/images/cover/:file", func(ctx *aero.Context) string {
		format := ".jpg"

		if ctx.CanUseWebP() {
			format = ".webp"
		}

		return ctx.File("images/cover/" + ctx.Get("file") + format)
	})

	app.Get("/images/elements/:file", func(ctx *aero.Context) string {
		return ctx.File("images/elements/" + ctx.Get("file"))
	})

	// For benchmarks
	app.Get("/hello", func(ctx *aero.Context) string {
		return ctx.Text("Hello World")
	})

	// Let's go
	app.Run()
}
