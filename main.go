package main

import (
	"errors"
	"net/http"

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
		return components.Layout(app, content)
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
		if ctx.CanUseWebP() {
			return ctx.File("images/icons/favicon.webp")
		}

		return ctx.File("images/icons/favicon.png")
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

	// Avatars
	app.Get("/user/:nick/avatar", func(ctx *aero.Context) string {
		nick := ctx.Get("nick")
		user, err := arn.GetUserByNick(nick)

		if err != nil {
			return ctx.Error(http.StatusNotFound, "User not found", err)
		}

		if ctx.CanUseWebP() {
			return ctx.File("images/avatars/large/webp/" + user.ID + ".webp")
		}

		original := arn.FindFileWithExtension(
			user.ID,
			"images/avatars/large/original/",
			arn.OriginalImageExtensions,
		)

		if original == "" {
			return ctx.Error(http.StatusNotFound, "Avatar not found", errors.New("Image not found for user: "+user.ID))
		}

		return ctx.File(original)
	})

	// Avatars
	app.Get("/user/:nick/avatar/small", func(ctx *aero.Context) string {
		nick := ctx.Get("nick")
		user, err := arn.GetUserByNick(nick)

		if err != nil {
			return ctx.Error(http.StatusNotFound, "User not found", err)
		}

		if ctx.CanUseWebP() {
			return ctx.File("images/avatars/small/webp/" + user.ID + ".webp")
		}

		original := arn.FindFileWithExtension(
			user.ID,
			"images/avatars/small/original/",
			arn.OriginalImageExtensions,
		)

		if original == "" {
			return ctx.Error(http.StatusNotFound, "Avatar not found", errors.New("Image not found for user: "+user.ID))
		}

		return ctx.File(original)
	})

	// Elements
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
