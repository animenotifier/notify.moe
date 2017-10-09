package main

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/session-store-aerospike"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/auth"
	"github.com/animenotifier/notify.moe/components/css"
	"github.com/animenotifier/notify.moe/layout"
	"github.com/animenotifier/notify.moe/middleware"
	"github.com/animenotifier/notify.moe/pages/admin"
	"github.com/animenotifier/notify.moe/pages/amvs"
	"github.com/animenotifier/notify.moe/pages/anime"
	"github.com/animenotifier/notify.moe/pages/animelist"
	"github.com/animenotifier/notify.moe/pages/animelistitem"
	"github.com/animenotifier/notify.moe/pages/apiview"
	"github.com/animenotifier/notify.moe/pages/artworks"
	"github.com/animenotifier/notify.moe/pages/best"
	"github.com/animenotifier/notify.moe/pages/character"
	"github.com/animenotifier/notify.moe/pages/charge"
	"github.com/animenotifier/notify.moe/pages/dashboard"
	"github.com/animenotifier/notify.moe/pages/editanime"
	"github.com/animenotifier/notify.moe/pages/editor"
	"github.com/animenotifier/notify.moe/pages/embed"
	"github.com/animenotifier/notify.moe/pages/explore"
	"github.com/animenotifier/notify.moe/pages/forum"
	"github.com/animenotifier/notify.moe/pages/forums"
	"github.com/animenotifier/notify.moe/pages/home"
	"github.com/animenotifier/notify.moe/pages/inventory"
	"github.com/animenotifier/notify.moe/pages/listimport"
	"github.com/animenotifier/notify.moe/pages/listimport/listimportanilist"
	"github.com/animenotifier/notify.moe/pages/listimport/listimportkitsu"
	"github.com/animenotifier/notify.moe/pages/listimport/listimportmyanimelist"
	"github.com/animenotifier/notify.moe/pages/login"
	"github.com/animenotifier/notify.moe/pages/me"
	"github.com/animenotifier/notify.moe/pages/newsoundtrack"
	"github.com/animenotifier/notify.moe/pages/newthread"
	"github.com/animenotifier/notify.moe/pages/notifications"
	"github.com/animenotifier/notify.moe/pages/paypal"
	"github.com/animenotifier/notify.moe/pages/posts"
	"github.com/animenotifier/notify.moe/pages/profile"
	"github.com/animenotifier/notify.moe/pages/search"
	"github.com/animenotifier/notify.moe/pages/settings"
	"github.com/animenotifier/notify.moe/pages/shop"
	"github.com/animenotifier/notify.moe/pages/soundtrack"
	"github.com/animenotifier/notify.moe/pages/soundtracks"
	"github.com/animenotifier/notify.moe/pages/statistics"
	"github.com/animenotifier/notify.moe/pages/threads"
	"github.com/animenotifier/notify.moe/pages/user"
	"github.com/animenotifier/notify.moe/pages/users"
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
	app.Sessions.Duration = 3600 * 24 * 30 * 6
	app.Sessions.Store = aerospikestore.New(arn.DB, "Session", app.Sessions.Duration)

	// Layout
	app.Layout = layout.Render

	// Ajax routes
	app.Ajax("/", home.Get)
	app.Ajax("/dashboard", dashboard.Get)
	app.Ajax("/anime/:id", anime.Get)
	app.Ajax("/anime/:id/edit", editanime.Get)
	app.Ajax("/api", apiview.Get)
	app.Ajax("/best/anime", best.Get)
	app.Ajax("/explore", explore.Get)
	app.Ajax("/forum", forums.Get)
	app.Ajax("/forum/:tag", forum.Get)
	app.Ajax("/thread/:id", threads.Get)
	app.Ajax("/post/:id", posts.Get)
	app.Ajax("/character/:id", character.Get)
	app.Ajax("/new/thread", newthread.Get)
	app.Ajax("/settings", settings.Get)
	app.Ajax("/artworks", artworks.Get)
	app.Ajax("/amvs", amvs.Get)
	app.Ajax("/users", users.Active)
	app.Ajax("/users/osu", users.Osu)
	app.Ajax("/users/staff", users.Staff)
	app.Ajax("/users/anime/watching", users.AnimeWatching)
	app.Ajax("/statistics", statistics.Get)
	app.Ajax("/statistics/anime", statistics.Anime)
	app.Ajax("/login", login.Get)

	// Soundtracks
	app.Ajax("/soundtracks", soundtracks.Get)
	app.Ajax("/new/soundtrack", newsoundtrack.Get)
	app.Ajax("/soundtrack/:id", soundtrack.Get)
	app.Ajax("/soundtrack/:id/edit", soundtrack.Edit)

	// User profiles
	app.Ajax("/user", user.Get)
	app.Ajax("/user/:nick", profile.Get)
	app.Ajax("/user/:nick/threads", profile.GetThreadsByUser)
	app.Ajax("/user/:nick/posts", profile.GetPostsByUser)
	app.Ajax("/user/:nick/soundtracks", profile.GetSoundTracksByUser)
	app.Ajax("/user/:nick/stats", profile.GetStatsByUser)
	app.Ajax("/user/:nick/followers", profile.GetFollowers)
	app.Ajax("/user/:nick/animelist", animelist.Get)
	app.Ajax("/user/:nick/animelist/watching", animelist.FilterByStatus(arn.AnimeListStatusWatching))
	app.Ajax("/user/:nick/animelist/completed", animelist.FilterByStatus(arn.AnimeListStatusCompleted))
	app.Ajax("/user/:nick/animelist/planned", animelist.FilterByStatus(arn.AnimeListStatusPlanned))
	app.Ajax("/user/:nick/animelist/hold", animelist.FilterByStatus(arn.AnimeListStatusHold))
	app.Ajax("/user/:nick/animelist/dropped", animelist.FilterByStatus(arn.AnimeListStatusDropped))
	app.Ajax("/user/:nick/animelist/anime/:id", animelistitem.Get)

	// Anime list
	app.Ajax("/animelist/watching", home.FilterByStatus(arn.AnimeListStatusWatching))
	app.Ajax("/animelist/completed", home.FilterByStatus(arn.AnimeListStatusCompleted))
	app.Ajax("/animelist/planned", home.FilterByStatus(arn.AnimeListStatusPlanned))
	app.Ajax("/animelist/hold", home.FilterByStatus(arn.AnimeListStatusHold))
	app.Ajax("/animelist/dropped", home.FilterByStatus(arn.AnimeListStatusDropped))

	// Search
	app.Ajax("/search", search.Get)
	app.Ajax("/search/:term", search.Get)

	// Shop
	app.Ajax("/shop", shop.Get)
	app.Ajax("/inventory", inventory.Get)
	app.Ajax("/charge", charge.Get)
	app.Ajax("/shop/history", shop.PurchaseHistory)
	app.Post("/api/shop/buy/:item/:quantity", shop.BuyItem)

	// Admin
	app.Ajax("/admin", admin.Get)
	app.Ajax("/admin/webdev", admin.WebDev)
	app.Ajax("/admin/purchases", admin.PurchaseHistory)

	// Editor
	app.Ajax("/editor", editor.Get)
	app.Ajax("/editor/anilist", editor.AniList)
	app.Ajax("/editor/shoboi", editor.Shoboi)

	// Import
	app.Ajax("/import", listimport.Get)
	app.Ajax("/import/anilist/animelist", listimportanilist.Preview)
	app.Ajax("/import/anilist/animelist/finish", listimportanilist.Finish)
	app.Ajax("/import/myanimelist/animelist", listimportmyanimelist.Preview)
	app.Ajax("/import/myanimelist/animelist/finish", listimportmyanimelist.Finish)
	app.Ajax("/import/kitsu/animelist", listimportkitsu.Preview)
	app.Ajax("/import/kitsu/animelist/finish", listimportkitsu.Finish)

	// Genres
	// app.Ajax("/genres", genres.Get)
	// app.Ajax("/genres/:name", genre.Get)

	// Browser extension
	app.Ajax("/extension/embed", embed.Get)

	// API
	app.Get("/api/me", me.Get)
	app.Get("/api/test/notification", notifications.Test)

	// PayPal
	app.Ajax("/paypal/success", paypal.Success)
	app.Ajax("/paypal/cancel", paypal.Cancel)
	app.Post("/api/paypal/payment/create", paypal.CreatePayment)

	// Assets
	configureAssets(app)

	// Rewrite
	app.Rewrite(rewrite)

	// Middleware
	app.Use(middleware.Firewall())
	app.Use(middleware.Log())
	app.Use(middleware.Session())
	app.Use(middleware.UserInfo())

	// API
	arn.API.Install(app)

	// Domain
	if arn.IsDevelopment() {
		app.Config.Domain = "beta.notify.moe"
	} else {
		arn.DB.SetScanPriority("high")
	}

	// Authentication
	auth.Install(app)

	// Specify test routes
	for route, examples := range routeTests {
		app.Test(route, examples)
	}

	return app
}
