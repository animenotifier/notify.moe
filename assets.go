package main

import (
	"io/ioutil"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components/js"
)

// configureAssets adds all the routes used for media assets.
func configureAssets(app *aero.Application) {
	// Script bundle
	scriptBundle := js.Bundle()

	// Service worker
	serviceWorkerBytes, err := ioutil.ReadFile("sw/service-worker.js")
	serviceWorker := string(serviceWorkerBytes)

	if err != nil {
		panic("Couldn't load service worker")
	}

	app.Get("/scripts", func(ctx *aero.Context) string {
		return ctx.JavaScript(scriptBundle)
	})

	app.Get("/scripts.js", func(ctx *aero.Context) string {
		return ctx.JavaScript(scriptBundle)
	})

	app.Get("/service-worker", func(ctx *aero.Context) string {
		return ctx.JavaScript(serviceWorker)
	})

	// Web manifest
	app.Get("/manifest.json", func(ctx *aero.Context) string {
		return ctx.JSON(app.Config.Manifest)
	})

	// Favicon
	app.Get("/favicon.ico", func(ctx *aero.Context) string {
		return ctx.TryWebP("images/brand/64", ".png")
	})

	// Images
	app.Get("/images/*file", func(ctx *aero.Context) string {
		return ctx.File("images" + ctx.Get("file"))
	})

	// For benchmarks
	app.Get("/hello", func(ctx *aero.Context) string {
		return ctx.Text("Hello World")
	})
}
