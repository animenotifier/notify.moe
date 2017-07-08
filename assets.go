package main

import (
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components/js"
)

func init() {
	// Scripts
	scripts := js.Bundle()

	app.Get("/scripts", func(ctx *aero.Context) string {
		ctx.SetResponseHeader("Content-Type", "application/javascript")
		return scripts
	})

	app.Get("/scripts.js", func(ctx *aero.Context) string {
		ctx.SetResponseHeader("Content-Type", "application/javascript")
		return scripts
	})

	// Web manifest
	app.Get("/manifest.json", func(ctx *aero.Context) string {
		return ctx.JSON(app.Config.Manifest)
	})

	// Favicon
	app.Get("/favicon.ico", func(ctx *aero.Context) string {
		return ctx.TryWebP("images/brand/64", ".png")
	})

	// Brand icons
	app.Get("/images/brand/:file", func(ctx *aero.Context) string {
		file := strings.TrimSuffix(ctx.Get("file"), ".webp")
		return ctx.TryWebP("images/brand/"+file, ".png")
	})

	// Cover image
	app.Get("/images/cover/:file", func(ctx *aero.Context) string {
		file := strings.TrimSuffix(ctx.Get("file"), ".webp")
		return ctx.TryWebP("images/cover/"+file, ".jpg")
	})

	// Login buttons
	app.Get("/images/login/:file", func(ctx *aero.Context) string {
		return ctx.File("images/login/" + ctx.Get("file") + ".png")
	})

	// Avatars
	app.Get("/images/avatars/large/:file", func(ctx *aero.Context) string {
		return ctx.File("images/avatars/large/" + ctx.Get("file"))
	})

	// Avatars
	app.Get("/images/avatars/small/:file", func(ctx *aero.Context) string {
		return ctx.File("images/avatars/large/" + ctx.Get("file"))
	})

	// Elements
	app.Get("/images/elements/:file", func(ctx *aero.Context) string {
		return ctx.File("images/elements/" + ctx.Get("file"))
	})

	// For benchmarks
	app.Get("/hello", func(ctx *aero.Context) string {
		return ctx.Text("Hello World")
	})
}
