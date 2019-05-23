package assets

import (
	"io/ioutil"
	"strings"

	"github.com/aerogo/aero"
	"github.com/aerogo/manifest"
	"github.com/aerogo/sitemap"
	"github.com/akyoto/stringutils/unsafe"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components/css"
	"github.com/animenotifier/notify.moe/components/js"
)

var (
	Manifest      *manifest.Manifest
	JS            string
	CSS           string
	ServiceWorker string
	Organization  string
)

// load loads all the necessary assets into memory.
func load() {
	var err error

	// Manifest
	Manifest, err = manifest.FromFile("manifest.json")

	if err != nil {
		panic("Couldn't load manifest.json")
	}

	// Service worker
	data, err := ioutil.ReadFile("scripts/ServiceWorker/ServiceWorker.js")

	if err != nil {
		panic("Couldn't load service worker")
	}

	ServiceWorker = unsafe.BytesToString(data)

	// Organization
	data, err = ioutil.ReadFile("organization.json")

	if err != nil {
		panic("Couldn't load organization.json")
	}

	Organization = unsafe.BytesToString(data)
	Organization = strings.ReplaceAll(Organization, "\n", "")
	Organization = strings.ReplaceAll(Organization, "\t", "")

	// Bundles
	JS = js.Bundle()
	CSS = css.Bundle()
}

// Configure adds all the routes used for media assets.
func Configure(app *aero.Application) {
	load()

	app.Get("/scripts", func(ctx *aero.Context) string {
		return ctx.JavaScript(JS)
	})

	app.Get("/styles", func(ctx *aero.Context) string {
		return ctx.CSS(CSS)
	})

	app.Get("/service-worker", func(ctx *aero.Context) string {
		return ctx.JavaScript(ServiceWorker)
	})

	// Web manifest
	app.Get("/manifest.json", func(ctx *aero.Context) string {
		return ctx.JSON(Manifest)
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
