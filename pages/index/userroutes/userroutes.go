package userroutes

import (
	"github.com/aerogo/layout"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/pages/animelist"
	"github.com/animenotifier/notify.moe/pages/animelistitem"
	"github.com/animenotifier/notify.moe/pages/compare"
	"github.com/animenotifier/notify.moe/pages/notifications"
	"github.com/animenotifier/notify.moe/pages/profile"
	"github.com/animenotifier/notify.moe/pages/profile/profilecharacters"
	"github.com/animenotifier/notify.moe/pages/profile/profilequotes"
	"github.com/animenotifier/notify.moe/pages/profile/profiletracks"
	"github.com/animenotifier/notify.moe/pages/recommended"
	"github.com/animenotifier/notify.moe/pages/user"
)

// Register registers the page routes.
func Register(l *layout.Layout) {
	// User profiles
	l.Page("/user", user.Get)
	l.Page("/user/:nick", profile.Get)
	l.Page("/user/:nick/characters/liked", profilecharacters.Liked)
	l.Page("/user/:nick/forum/threads", profile.GetThreadsByUser)
	l.Page("/user/:nick/forum/posts", profile.GetPostsByUser)
	l.Page("/user/:nick/soundtracks/added", profiletracks.Added)
	l.Page("/user/:nick/soundtracks/added/from/:index", profiletracks.Added)
	l.Page("/user/:nick/soundtracks/liked", profiletracks.Liked)
	l.Page("/user/:nick/soundtracks/liked/from/:index", profiletracks.Liked)
	l.Page("/user/:nick/quotes/added", profilequotes.Added)
	l.Page("/user/:nick/quotes/added/from/:index", profilequotes.Added)
	l.Page("/user/:nick/quotes/liked", profilequotes.Liked)
	l.Page("/user/:nick/quotes/liked/from/:index", profilequotes.Liked)
	l.Page("/user/:nick/stats", profile.GetStatsByUser)
	l.Page("/user/:nick/followers", profile.GetFollowers)
	l.Page("/user/:nick/animelist/anime/:id", animelistitem.Get)
	l.Page("/user/:nick/recommended/anime", recommended.Anime)
	l.Page("/user/:nick/notifications", notifications.ByUser)
	l.Page("/user/:nick/edit", user.Edit)

	// Anime list
	l.Page("/user/:nick/animelist/watching", animelist.FilterByStatus(arn.AnimeListStatusWatching))
	l.Page("/user/:nick/animelist/completed", animelist.FilterByStatus(arn.AnimeListStatusCompleted))
	l.Page("/user/:nick/animelist/planned", animelist.FilterByStatus(arn.AnimeListStatusPlanned))
	l.Page("/user/:nick/animelist/hold", animelist.FilterByStatus(arn.AnimeListStatusHold))
	l.Page("/user/:nick/animelist/dropped", animelist.FilterByStatus(arn.AnimeListStatusDropped))

	l.Page("/user/:nick/animelist/watching/from/:index", animelist.FilterByStatus(arn.AnimeListStatusWatching))
	l.Page("/user/:nick/animelist/completed/from/:index", animelist.FilterByStatus(arn.AnimeListStatusCompleted))
	l.Page("/user/:nick/animelist/planned/from/:index", animelist.FilterByStatus(arn.AnimeListStatusPlanned))
	l.Page("/user/:nick/animelist/hold/from/:index", animelist.FilterByStatus(arn.AnimeListStatusHold))
	l.Page("/user/:nick/animelist/dropped/from/:index", animelist.FilterByStatus(arn.AnimeListStatusDropped))

	// Redirects
	l.Page("/animelist/watching", animelist.Redirect)
	l.Page("/animelist/completed", animelist.Redirect)
	l.Page("/animelist/planned", animelist.Redirect)
	l.Page("/animelist/hold", animelist.Redirect)
	l.Page("/animelist/dropped", animelist.Redirect)

	// Compare
	l.Page("/compare/animelist/:nick-1/:nick-2", compare.AnimeList)

	// Notifications
	l.Page("/notifications", notifications.ByUser)
	l.Page("/notifications/all", notifications.All)
}
