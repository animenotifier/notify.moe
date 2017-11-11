package pages

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/layout"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/layout"
	"github.com/animenotifier/notify.moe/pages/admin"
	"github.com/animenotifier/notify.moe/pages/anime"
	"github.com/animenotifier/notify.moe/pages/animelist"
	"github.com/animenotifier/notify.moe/pages/animelistitem"
	"github.com/animenotifier/notify.moe/pages/apiview"
	"github.com/animenotifier/notify.moe/pages/character"
	"github.com/animenotifier/notify.moe/pages/charge"
	"github.com/animenotifier/notify.moe/pages/compare"
	"github.com/animenotifier/notify.moe/pages/database"
	"github.com/animenotifier/notify.moe/pages/editanime"
	"github.com/animenotifier/notify.moe/pages/editor"
	"github.com/animenotifier/notify.moe/pages/embed"
	"github.com/animenotifier/notify.moe/pages/explore"
	"github.com/animenotifier/notify.moe/pages/forum"
	"github.com/animenotifier/notify.moe/pages/forums"
	"github.com/animenotifier/notify.moe/pages/group"
	"github.com/animenotifier/notify.moe/pages/groups"
	"github.com/animenotifier/notify.moe/pages/home"
	"github.com/animenotifier/notify.moe/pages/inventory"
	"github.com/animenotifier/notify.moe/pages/listimport"
	"github.com/animenotifier/notify.moe/pages/listimport/listimportanilist"
	"github.com/animenotifier/notify.moe/pages/listimport/listimportkitsu"
	"github.com/animenotifier/notify.moe/pages/listimport/listimportmyanimelist"
	"github.com/animenotifier/notify.moe/pages/login"
	"github.com/animenotifier/notify.moe/pages/me"
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

// Configure registers the page routes in the application.
func Configure(app *aero.Application) {
	layout := layout.New(app)

	// Set render function for the layout
	layout.Render = fullpage.Render

	// Main menu
	layout.Page("/", home.Get)
	layout.Page("/explore", explore.Get)
	layout.Page("/explore/anime/:year/:status/:type", explore.Filter)
	layout.Page("/login", login.Get)
	layout.Page("/api", apiview.Get)
	// layout.Ajax("/dashboard", dashboard.Get)
	// layout.Ajax("/best/anime", best.Get)
	// layout.Ajax("/artworks", artworks.Get)
	// layout.Ajax("/amvs", amvs.Get)

	// Forum
	layout.Page("/forum", forums.Get)
	layout.Page("/forum/:tag", forum.Get)
	layout.Page("/thread/:id", threads.Get)
	layout.Page("/post/:id", posts.Get)
	layout.Page("/new/thread", newthread.Get)

	// User lists
	layout.Page("/users", users.Active)
	layout.Page("/users/osu", users.Osu)
	layout.Page("/users/staff", users.Staff)

	// Statistics
	layout.Page("/statistics", statistics.Get)
	layout.Page("/statistics/anime", statistics.Anime)

	// Anime
	layout.Page("/anime/:id", anime.Get)
	layout.Page("/anime/:id/episodes", anime.Episodes)
	layout.Page("/anime/:id/characters", anime.Characters)
	layout.Page("/anime/:id/tracks", anime.Tracks)
	layout.Page("/anime/:id/edit", editanime.Get)

	// Characters
	layout.Page("/character/:id", character.Get)

	// Settings
	layout.Page("/settings", settings.Get(components.SettingsPersonal))
	layout.Page("/settings/accounts", settings.Get(components.SettingsAccounts))
	layout.Page("/settings/notifications", settings.Get(components.SettingsNotifications))
	layout.Page("/settings/apps", settings.Get(components.SettingsApps))
	layout.Page("/settings/avatar", settings.Get(components.SettingsAvatar))
	layout.Page("/settings/formatting", settings.Get(components.SettingsFormatting))
	layout.Page("/settings/pro", settings.Get(components.SettingsPro))

	// Soundtracks
	layout.Page("/soundtracks", soundtracks.Get)
	layout.Page("/soundtracks/from/:index", soundtracks.From)
	layout.Page("/soundtrack/:id", soundtrack.Get)
	layout.Page("/soundtrack/:id/edit", soundtrack.Edit)

	// Groups
	layout.Page("/groups", groups.Get)
	layout.Page("/group/:id", group.Get)
	layout.Page("/group/:id/edit", group.Edit)
	layout.Page("/group/:id/forum", group.Forum)

	// User profiles
	layout.Page("/user", user.Get)
	layout.Page("/user/:nick", profile.Get)
	layout.Page("/user/:nick/threads", profile.GetThreadsByUser)
	layout.Page("/user/:nick/posts", profile.GetPostsByUser)
	layout.Page("/user/:nick/soundtracks", profile.GetSoundTracksByUser)
	layout.Page("/user/:nick/stats", profile.GetStatsByUser)
	layout.Page("/user/:nick/followers", profile.GetFollowers)
	layout.Page("/user/:nick/animelist", animelist.Get)
	layout.Page("/user/:nick/animelist/watching", animelist.FilterByStatus(arn.AnimeListStatusWatching))
	layout.Page("/user/:nick/animelist/completed", animelist.FilterByStatus(arn.AnimeListStatusCompleted))
	layout.Page("/user/:nick/animelist/planned", animelist.FilterByStatus(arn.AnimeListStatusPlanned))
	layout.Page("/user/:nick/animelist/hold", animelist.FilterByStatus(arn.AnimeListStatusHold))
	layout.Page("/user/:nick/animelist/dropped", animelist.FilterByStatus(arn.AnimeListStatusDropped))
	layout.Page("/user/:nick/animelist/anime/:id", animelistitem.Get)

	// Anime list
	layout.Page("/animelist/watching", home.FilterByStatus(arn.AnimeListStatusWatching))
	layout.Page("/animelist/completed", home.FilterByStatus(arn.AnimeListStatusCompleted))
	layout.Page("/animelist/planned", home.FilterByStatus(arn.AnimeListStatusPlanned))
	layout.Page("/animelist/hold", home.FilterByStatus(arn.AnimeListStatusHold))
	layout.Page("/animelist/dropped", home.FilterByStatus(arn.AnimeListStatusDropped))

	// Compare
	layout.Page("/compare/animelist/:nick-1/:nick-2", compare.AnimeList)

	// Search
	layout.Page("/search", search.Get)
	layout.Page("/search/:term", search.Get)

	// Shop
	layout.Page("/shop", shop.Get)
	layout.Page("/inventory", inventory.Get)
	layout.Page("/charge", charge.Get)
	layout.Page("/shop/history", shop.PurchaseHistory)
	app.Post("/api/shop/buy/:item/:quantity", shop.BuyItem)

	// Admin
	layout.Page("/admin", admin.Get)
	layout.Page("/admin/webdev", admin.WebDev)
	layout.Page("/admin/purchases", admin.PurchaseHistory)

	// Editor
	layout.Page("/editor", editor.Get)
	layout.Page("/editor/anilist", editor.AniList)
	layout.Page("/editor/shoboi", editor.Shoboi)

	// Mixed
	layout.Page("/database", database.Get)
	app.Get("/api/select/:data-type/where/:field/is/:field-value", database.Select)

	// Import
	layout.Page("/import", listimport.Get)
	layout.Page("/import/anilist/animelist", listimportanilist.Preview)
	layout.Page("/import/anilist/animelist/finish", listimportanilist.Finish)
	layout.Page("/import/myanimelist/animelist", listimportmyanimelist.Preview)
	layout.Page("/import/myanimelist/animelist/finish", listimportmyanimelist.Finish)
	layout.Page("/import/kitsu/animelist", listimportkitsu.Preview)
	layout.Page("/import/kitsu/animelist/finish", listimportkitsu.Finish)

	// Browser extension
	layout.Page("/extension/embed", embed.Get)

	// API
	app.Get("/api/me", me.Get)
	app.Get("/api/test/notification", notifications.Test)

	// PayPal
	layout.Page("/paypal/success", paypal.Success)
	layout.Page("/paypal/cancel", paypal.Cancel)
	app.Post("/api/paypal/payment/create", paypal.CreatePayment)

	// Genres
	// layout.Ajax("/genres", genres.Get)
	// layout.Ajax("/genres/:name", genre.Get)
}
