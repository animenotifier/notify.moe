package pages

import (
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/pages/amvs"
	"github.com/animenotifier/notify.moe/pages/anime/editanime"
	"github.com/animenotifier/notify.moe/pages/animeimport"
	"github.com/animenotifier/notify.moe/pages/animelist"
	"github.com/animenotifier/notify.moe/pages/animelistitem"
	"github.com/animenotifier/notify.moe/pages/apiview"
	"github.com/animenotifier/notify.moe/pages/apiview/apidocs"
	"github.com/animenotifier/notify.moe/pages/calendar"
	"github.com/animenotifier/notify.moe/pages/characters"
	"github.com/animenotifier/notify.moe/pages/compare"
	"github.com/animenotifier/notify.moe/pages/database"
	"github.com/animenotifier/notify.moe/pages/editor/jobs"
	"github.com/animenotifier/notify.moe/pages/embed"
	"github.com/animenotifier/notify.moe/pages/episode"
	"github.com/animenotifier/notify.moe/pages/explore/explorecolor"
	"github.com/animenotifier/notify.moe/pages/explore/explorerelations"
	"github.com/animenotifier/notify.moe/pages/explore/halloffame"
	"github.com/animenotifier/notify.moe/pages/genre"
	"github.com/animenotifier/notify.moe/pages/genres"
	"github.com/animenotifier/notify.moe/pages/home"
	"github.com/animenotifier/notify.moe/pages/index/companyroutes"
	"github.com/animenotifier/notify.moe/pages/index/forumroutes"
	"github.com/animenotifier/notify.moe/pages/index/grouproutes"
	"github.com/animenotifier/notify.moe/pages/index/importroutes"
	"github.com/animenotifier/notify.moe/pages/index/quoteroutes"
	"github.com/animenotifier/notify.moe/pages/index/searchroutes"
	"github.com/animenotifier/notify.moe/pages/index/settingsroutes"
	"github.com/animenotifier/notify.moe/pages/index/shoproutes"
	"github.com/animenotifier/notify.moe/pages/index/soundtrackroutes"
	"github.com/animenotifier/notify.moe/pages/index/staffroutes"
	"github.com/animenotifier/notify.moe/pages/index/userlistroutes"
	"github.com/animenotifier/notify.moe/pages/login"
	"github.com/animenotifier/notify.moe/pages/me"
	"github.com/animenotifier/notify.moe/pages/notifications"
	"github.com/animenotifier/notify.moe/pages/popular"
	"github.com/animenotifier/notify.moe/pages/post"
	"github.com/animenotifier/notify.moe/pages/profile"
	"github.com/animenotifier/notify.moe/pages/profile/profilecharacters"
	"github.com/animenotifier/notify.moe/pages/recommended"
	"github.com/animenotifier/notify.moe/pages/soundtrack"
	"github.com/animenotifier/notify.moe/pages/sse"
	"github.com/animenotifier/notify.moe/pages/statistics"
	"github.com/animenotifier/notify.moe/pages/terms"
	"github.com/animenotifier/notify.moe/pages/thread"
	"github.com/animenotifier/notify.moe/pages/upload"
	"github.com/animenotifier/notify.moe/pages/welcome"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Configure registers the page routes in the application.
func Configure(app *aero.Application) {
	core(app)
	activity(app)
	anime(app)
	api(app)
	user(app)
	character(app)
	explore(app)
	amv(app)

	forumroutes.Register(app)
	userlistroutes.Register(app)
	quoteroutes.Register(app)
	companyroutes.Register(app)
	soundtrackroutes.Register(app)
	grouproutes.Register(app)
	searchroutes.Register(app)
	importroutes.Register(app)
	shoproutes.Register(app)
	settingsroutes.Register(app)
	staffroutes.Register(app)
}

func amv(app *aero.Application) {
	// AMV
	page.Get(app, "/amv/:id", amv.Get)
	page.Get(app, "/amv/:id/edit", amv.Edit)
	page.Get(app, "/amv/:id/history", amv.History)

	// AMVs
	page.Get(app, "/amvs", amvs.Latest)
	page.Get(app, "/amvs/from/:index", amvs.Latest)
	page.Get(app, "/amvs/best", amvs.Best)
	page.Get(app, "/amvs/best/from/:index", amvs.Best)
}

func explore(app *aero.Application) {
	// Explore
	page.Get(app, "/explore", explore.Filter)
	page.Get(app, "/explore/anime/:year/:season/:status/:type", explore.Filter)
	page.Get(app, "/explore/color/:color/anime", explorecolor.AnimeByAverageColor)
	page.Get(app, "/explore/color/:color/anime/from/:index", explorecolor.AnimeByAverageColor)
	page.Get(app, "/halloffame", halloffame.Get)
}

func character(app *aero.Application) {
	// Character
	page.Get(app, "/character/:id", character.Get)
	page.Get(app, "/character/:id/edit", character.Edit)
	page.Get(app, "/character/:id/edit/images", character.EditImages)
	page.Get(app, "/character/:id/history", character.History)

	// Characters
	page.Get(app, "/characters", characters.Latest)
	page.Get(app, "/characters/from/:index", characters.Latest)
	page.Get(app, "/characters/best", characters.Best)
	page.Get(app, "/characters/best/from/:index", characters.Best)
}

func core(app *aero.Application) {
	// Core
	page.Get(app, "/", home.Get)
	page.Get(app, "/login", login.Get)
	page.Get(app, "/welcome", welcome.Get)
	page.Get(app, "/terms", terms.Get)
	page.Get(app, "/extension/embed", embed.Get)

	// Calendar
	page.Get(app, "/calendar", calendar.Get)

	// Statistics
	page.Get(app, "/statistics", statistics.Get)
	page.Get(app, "/statistics/anime", statistics.Anime)
}

func activity(app *aero.Application) {
	// Activity
	page.Get(app, "/activity", activity.Global)
	page.Get(app, "/activity/from/:index", activity.Global)
	page.Get(app, "/activity/followed", activity.Followed)
	page.Get(app, "/activity/followed/from/:index", activity.Followed)
}

func anime(app *aero.Application) {
	// Anime
	page.Get(app, "/anime/:id", anime.Get)
	page.Get(app, "/anime/:id/episodes", anime.Episodes)
	page.Get(app, "/anime/:id/characters", anime.Characters)
	page.Get(app, "/anime/:id/tracks", anime.Tracks)
	page.Get(app, "/anime/:id/relations", anime.Relations)
	page.Get(app, "/anime/:id/comments", anime.Comments)
	page.Get(app, "/episode/:id", episode.Get)
	app.Get("/episode/:id/subtitles/:language", episode.Subtitles)

	// Editing
	page.Get(app, "/anime/:id/edit", editanime.Main)
	page.Get(app, "/anime/:id/edit/images", editanime.Images)
	page.Get(app, "/anime/:id/edit/characters", editanime.Characters)
	page.Get(app, "/anime/:id/edit/relations", editanime.Relations)
	page.Get(app, "/anime/:id/edit/episodes", editanime.Episodes)
	page.Get(app, "/anime/:id/edit/history", editanime.History)

	// Redirects
	page.Get(app, "/kitsu/anime/:id", anime.RedirectByMapping("kitsu/anime"))
	page.Get(app, "/mal/anime/:id", anime.RedirectByMapping("myanimelist/anime"))
	page.Get(app, "/anilist/anime/:id", anime.RedirectByMapping("anilist/anime"))

	// Genres
	page.Get(app, "/genres", genres.Get)
	page.Get(app, "/genre/:name", genre.Get)
}

func user(app *aero.Application) {
	// User
	page.Get(app, "/user", user.Get)
	page.Get(app, "/user/:nick", profile.Get)
	page.Get(app, "/user/:nick/characters/liked", profilecharacters.Liked)
	page.Get(app, "/user/:nick/animelist/anime/:id", animelistitem.Get)
	page.Get(app, "/user/:nick/anime/recommended", recommended.Anime)
	page.Get(app, "/user/:nick/anime/sequels", explorerelations.Sequels)
	page.Get(app, "/user/:nick/notifications", notifications.ByUser)
	page.Get(app, "/user/:nick/edit", user.Edit)
	// page.Get(app, "/user/:nick/forum/threads", profile.GetThreadsByUser)
	// page.Get(app, "/user/:nick/forum/posts", profile.GetPostsByUser)
	// page.Get(app, "/user/:nick/soundtracks/added", profiletracks.Added)
	// page.Get(app, "/user/:nick/soundtracks/added/from/:index", profiletracks.Added)
	// page.Get(app, "/user/:nick/soundtracks/liked", profiletracks.Liked)
	// page.Get(app, "/user/:nick/soundtracks/liked/from/:index", profiletracks.Liked)
	// page.Get(app, "/user/:nick/quotes/added", profilequotes.Added)
	// page.Get(app, "/user/:nick/quotes/added/from/:index", profilequotes.Added)
	// page.Get(app, "/user/:nick/quotes/liked", profilequotes.Liked)
	// page.Get(app, "/user/:nick/quotes/liked/from/:index", profilequotes.Liked)
	// page.Get(app, "/user/:nick/stats", profile.GetStatsByUser)
	// page.Get(app, "/user/:nick/followers", profile.GetFollowers)

	// Anime list
	page.Get(app, "/user/:nick/animelist/:status", animelist.Filter)
	page.Get(app, "/user/:nick/animelist/:status/from/:index", animelist.Filter)
	page.Get(app, "/compare/animelist/:nick-1/:nick-2", compare.AnimeList)
	page.Get(app, "/animelist/delete", animelist.DeleteConfirmation)

	// Redirects
	page.Get(app, "/animelist/watching", animelist.Redirect)
	page.Get(app, "/animelist/completed", animelist.Redirect)
	page.Get(app, "/animelist/planned", animelist.Redirect)
	page.Get(app, "/animelist/hold", animelist.Redirect)
	page.Get(app, "/animelist/dropped", animelist.Redirect)

	// Notifications
	page.Get(app, "/notifications", notifications.ByUser)
	page.Get(app, "/notifications/all", notifications.All)
}

func api(app *aero.Application) {
	// API documentation
	page.Get(app, "/api", apiview.Get)

	for name := range arn.DB.Types() {
		page.Get(app, "/api/"+strings.ToLower(name), apidocs.ByType(name))
	}

	// API
	app.Get("/api/me", me.Get)
	app.Get("/api/popular/anime/titles/:count", popular.AnimeTitles)
	app.Get("/api/test/notification", notifications.Test)
	app.Get("/api/count/notifications/unseen", notifications.CountUnseen)
	app.Get("/api/mark/notifications/seen", notifications.MarkNotificationsAsSeen)
	app.Get("/api/user/:id/notifications/latest", notifications.Latest)
	app.Get("/api/random/soundtrack", soundtrack.Random)
	app.Get("/api/next/soundtrack", soundtrack.Next)
	app.Get("/api/character/:id/ranking", character.Ranking)

	// Live updates
	app.Get("/api/sse/events", sse.Events)

	// Thread
	app.Get("/api/thread/:id/reply/ui", thread.ReplyUI)

	// Post
	app.Get("/api/post/:id/reply/ui", post.ReplyUI)

	// SoundTrack
	app.Post("/api/soundtrack/:id/download", soundtrack.Download)

	// AnimeList
	app.Post("/api/delete/animelist", animelist.Delete)

	// Upload
	app.Post("/api/upload/user/image", upload.UserImage)
	app.Post("/api/upload/user/cover", upload.UserCover)
	app.Post("/api/upload/anime/:id/image", upload.AnimeImage)
	app.Post("/api/upload/character/:id/image", upload.CharacterImage)
	app.Post("/api/upload/group/:id/image", upload.GroupImage)
	app.Post("/api/upload/amv/:id/file", upload.AMVFile)

	// Import anime
	app.Post("/api/import/kitsu/anime/:id", animeimport.Kitsu)
	app.Post("/api/delete/kitsu/anime/:id", animeimport.DeleteKitsu)

	// Jobs
	app.Post("/api/job/:job/start", jobs.Start)

	// Database
	app.Get("/api/types", database.Types)
}
