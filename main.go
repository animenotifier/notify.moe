package main

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/api"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/auth"
	"github.com/animenotifier/notify.moe/components/css"
	"github.com/animenotifier/notify.moe/layout"
	"github.com/animenotifier/notify.moe/middleware"
	"github.com/animenotifier/notify.moe/pages/admin"
	"github.com/animenotifier/notify.moe/pages/anime"
	"github.com/animenotifier/notify.moe/pages/animelist"
	"github.com/animenotifier/notify.moe/pages/animelistitem"
	"github.com/animenotifier/notify.moe/pages/apiview"
	"github.com/animenotifier/notify.moe/pages/best"
	"github.com/animenotifier/notify.moe/pages/dashboard"
	"github.com/animenotifier/notify.moe/pages/editanime"
	"github.com/animenotifier/notify.moe/pages/embed"
	"github.com/animenotifier/notify.moe/pages/explore"
	"github.com/animenotifier/notify.moe/pages/forum"
	"github.com/animenotifier/notify.moe/pages/forums"
	"github.com/animenotifier/notify.moe/pages/login"
	"github.com/animenotifier/notify.moe/pages/music"
	"github.com/animenotifier/notify.moe/pages/newsoundtrack"
	"github.com/animenotifier/notify.moe/pages/newthread"
	"github.com/animenotifier/notify.moe/pages/posts"
	"github.com/animenotifier/notify.moe/pages/profile"
	"github.com/animenotifier/notify.moe/pages/search"
	"github.com/animenotifier/notify.moe/pages/settings"
	"github.com/animenotifier/notify.moe/pages/threads"
	"github.com/animenotifier/notify.moe/pages/tracks"
	"github.com/animenotifier/notify.moe/pages/user"
	"github.com/animenotifier/notify.moe/pages/users"
	"github.com/animenotifier/notify.moe/pages/webdev"
)

var app = aero.New()

func main() {
	// Configure and start
	configure(app).Run()
}

func configure(app *aero.Application) *aero.Application {
	// HTTPS
	app.Security.Load("security/fullchain.pem", "security/privkey.pem")

	// CSS
	app.SetStyle(css.Bundle())

	// Sessions
	app.Sessions.Duration = 3600 * 24 * 7
	app.Sessions.Store = arn.NewAerospikeStore("Session", app.Sessions.Duration)

	// Layout
	app.Layout = layout.Render

	// Ajax routes
	app.Ajax("/", dashboard.Get)
	app.Ajax("/anime/:id", anime.Get)
	app.Ajax("/anime/:id/edit", editanime.Get)
	app.Ajax("/api", apiview.Get)
	app.Ajax("/best/anime", best.Get)
	app.Ajax("/explore", explore.Get)
	app.Ajax("/forum", forums.Get)
	app.Ajax("/forum/:tag", forum.Get)
	app.Ajax("/threads/:id", threads.Get)
	app.Ajax("/posts/:id", posts.Get)
	app.Ajax("/tracks/:id", tracks.Get)
	app.Ajax("/user", user.Get)
	app.Ajax("/user/:nick", profile.Get)
	app.Ajax("/user/:nick/threads", profile.GetThreadsByUser)
	app.Ajax("/user/:nick/posts", profile.GetPostsByUser)
	app.Ajax("/user/:nick/tracks", profile.GetSoundTracksByUser)
	app.Ajax("/user/:nick/animelist", animelist.Get)
	app.Ajax("/user/:nick/animelist/:id", animelistitem.Get)
	app.Ajax("/new/thread", newthread.Get)
	app.Ajax("/new/soundtrack", newsoundtrack.Get)
	app.Ajax("/settings", settings.Get)
	app.Ajax("/music", music.Get)
	app.Ajax("/admin", admin.Get)
	app.Ajax("/search", search.Get)
	app.Ajax("/search/:term", search.Get)
	app.Ajax("/users", users.Get)
	app.Ajax("/login", login.Get)
	app.Ajax("/webdev", webdev.Get)
	app.Ajax("/extension/embed", embed.Get)
	// app.Ajax("/genres", genres.Get)
	// app.Ajax("/genres/:name", genre.Get)

	// Middleware
	app.Use(middleware.Log())
	app.Use(middleware.Session())
	app.Use(middleware.UserInfo())

	// API
	api := api.New("/api/", arn.DB)
	api.Install(app)

	// Domain
	if arn.IsDevelopment() {
		app.Config.Domain = "beta.notify.moe"
	} else {
		arn.DB.SetScanPriority("high")
	}

	// Authentication
	auth.Install(app)

	return app
}
