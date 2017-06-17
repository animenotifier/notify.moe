package main

import (
	"errors"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

func init() {
	// Favicon
	app.Get("/favicon.ico", func(ctx *aero.Context) string {
		return ctx.TryWebP("images/brand/64", ".png")
	})

	// Favicon
	app.Get("/images/brand/:size", func(ctx *aero.Context) string {
		return ctx.TryWebP("images/brand/"+ctx.Get("size"), ".png")
	})

	// Scripts
	app.Get("/scripts.js", func(ctx *aero.Context) string {
		return ctx.File("temp/scripts.js")
	})

	// Web manifest
	app.Get("/manifest.json", func(ctx *aero.Context) string {
		return ctx.JSON(app.Config.Manifest)
	})

	// SVG icons
	app.Get("/icons/:file", func(ctx *aero.Context) string {
		return ctx.File("images/icons/svg/" + ctx.Get("file") + ".svg")
	})

	// Cover image
	app.Get("/images/cover/:file", func(ctx *aero.Context) string {
		return ctx.TryWebP("images/cover/"+ctx.Get("file"), ".jpg")
	})

	// Login buttons
	app.Get("/images/login/:file", func(ctx *aero.Context) string {
		return ctx.File("images/login/" + ctx.Get("file") + ".png")
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
}
