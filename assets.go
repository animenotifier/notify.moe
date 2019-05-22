package main

import (
	"io/ioutil"

	"github.com/aerogo/aero"
	"github.com/aerogo/manifest"
	"github.com/aerogo/sitemap"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components/css"
	"github.com/animenotifier/notify.moe/components/js"
)

// configureAssets adds all the routes used for media assets.
func configureAssets(app *aero.Application) {
	// Script bundle
	scriptBundle := js.Bundle()

	// Service worker
	serviceWorkerBytes, err := ioutil.ReadFile("scripts/ServiceWorker/ServiceWorker.js")

	if err != nil {
		panic("Couldn't load service worker")
	}

	serviceWorker := string(serviceWorkerBytes)

	// CSS bundle
	cssBundle := css.Bundle()

	// Manifest
	webManifest, err := manifest.FromFile("manifest.json")

	if err != nil {
		panic("Couldn't load web manifest")
	}

	app.Get("/scripts", func(ctx *aero.Context) string {
		return ctx.JavaScript(scriptBundle)
	})

	app.Get("/styles", func(ctx *aero.Context) string {
		return ctx.CSS(cssBundle)
	})

	app.Get("/service-worker", func(ctx *aero.Context) string {
		return ctx.JavaScript(serviceWorker)
	})

	// Web manifest
	app.Get("/manifest.json", func(ctx *aero.Context) string {
		return ctx.JSON(webManifest)
	})

	// Favicon
	app.Get("/favicon.ico", func(ctx *aero.Context) string {
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return ctx.File("images/brand/64.png")
	})

	// Images
	app.Get("/images/*file", func(ctx *aero.Context) string {
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return ctx.File("images/" + ctx.Get("file"))
	})

	// Videos
	app.Get("/videos/*file", func(ctx *aero.Context) string {
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return ctx.File("videos/" + ctx.Get("file"))
	})

	// Audio
	app.Get("/audio/*file", func(ctx *aero.Context) string {
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return ctx.File("audio/" + ctx.Get("file"))
	})

	// Anime sitemap
	app.Get("/sitemap/anime.txt", func(ctx *aero.Context) string {
		sitemap := sitemap.New()
		prefix := "https://" + app.Config.Domain

		for anime := range arn.StreamAnime() {
			sitemap.Add(prefix + anime.Link())
		}

		return ctx.Text(sitemap.Text())
	})

	// Character sitemap
	app.Get("/sitemap/character.txt", func(ctx *aero.Context) string {
		sitemap := sitemap.New()
		prefix := "https://" + app.Config.Domain

		for character := range arn.StreamCharacters() {
			sitemap.Add(prefix + character.Link())
		}

		return ctx.Text(sitemap.Text())
	})

	// User sitemap
	app.Get("/sitemap/user.txt", func(ctx *aero.Context) string {
		sitemap := sitemap.New()
		prefix := "https://" + app.Config.Domain

		for user := range arn.StreamUsers() {
			if !user.HasNick() {
				continue
			}

			sitemap.Add(prefix + user.Link())
		}

		return ctx.Text(sitemap.Text())
	})

	// SoundTrack sitemap
	app.Get("/sitemap/soundtrack.txt", func(ctx *aero.Context) string {
		sitemap := sitemap.New()
		prefix := "https://" + app.Config.Domain

		for soundTrack := range arn.StreamSoundTracks() {
			sitemap.Add(prefix + soundTrack.Link())
		}

		return ctx.Text(sitemap.Text())
	})

	// Thread sitemap
	app.Get("/sitemap/thread.txt", func(ctx *aero.Context) string {
		sitemap := sitemap.New()
		prefix := "https://" + app.Config.Domain

		for thread := range arn.StreamThreads() {
			sitemap.Add(prefix + thread.Link())
		}

		return ctx.Text(sitemap.Text())
	})

	// Post sitemap
	app.Get("/sitemap/post.txt", func(ctx *aero.Context) string {
		sitemap := sitemap.New()
		prefix := "https://" + app.Config.Domain

		for post := range arn.StreamPosts() {
			sitemap.Add(prefix + post.Link())
		}

		return ctx.Text(sitemap.Text())
	})

	// For benchmarks
	app.Get("/hello", func(ctx *aero.Context) string {
		return ctx.Text("Hello World")
	})
}
