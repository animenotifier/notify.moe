package staffroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/admin"
	"github.com/animenotifier/notify.moe/pages/editlog"
	"github.com/animenotifier/notify.moe/pages/editor"
	"github.com/animenotifier/notify.moe/pages/editor/filteranime"
	"github.com/animenotifier/notify.moe/pages/editor/filtercompanies"
	"github.com/animenotifier/notify.moe/pages/editor/filtersoundtracks"
	"github.com/animenotifier/notify.moe/pages/editor/jobs"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	// Editor
	page.Get(app, "/editor", editor.Get)

	// Editor links can be filtered by year, status and type
	editorFilterable := func(route string, handler func(ctx aero.Context) error) {
		page.Get(app, route+"/:year/:season/:status/:type", handler)
	}

	// Editor - Anime
	editorFilterable("/editor/anime/mapping/shoboi", filteranime.Shoboi)
	editorFilterable("/editor/anime/mapping/anilist", filteranime.AniList)
	editorFilterable("/editor/anime/mapping/mal", filteranime.MAL)
	editorFilterable("/editor/anime/mapping/duplicate", filteranime.DuplicateMappings)

	editorFilterable("/editor/anime/image/lowres", filteranime.LowResolutionAnimeImages)
	editorFilterable("/editor/anime/image/ultralowres", filteranime.UltraLowResolutionAnimeImages)

	editorFilterable("/editor/anime/companies/studios", filteranime.Studios)
	editorFilterable("/editor/anime/companies/producers", filteranime.Producers)
	editorFilterable("/editor/anime/companies/licensors", filteranime.Licensors)

	editorFilterable("/editor/anime/connections/relations", filteranime.Relations)
	editorFilterable("/editor/anime/connections/characters", filteranime.Characters)

	editorFilterable("/editor/anime/details/synopsis", filteranime.Synopsis)
	editorFilterable("/editor/anime/details/genres", filteranime.Genres)
	editorFilterable("/editor/anime/details/trailers", filteranime.Trailers)
	editorFilterable("/editor/anime/details/startdate", filteranime.StartDate)
	editorFilterable("/editor/anime/details/episodelength", filteranime.EpisodeLength)
	editorFilterable("/editor/anime/details/source", filteranime.Source)

	editorFilterable("/editor/anime/all", filteranime.All)

	// Editor - MALdiff
	editorFilterable("/editor/mal/diff/anime", editor.CompareMAL)

	// Editor - Kitsu
	page.Get(app, "/editor/kitsu/new/anime", editor.NewKitsuAnime)

	// Editor - Companies
	page.Get(app, "/editor/companies/description", filtercompanies.NoDescription)

	// Editor - Soundtracks
	page.Get(app, "/editor/soundtracks/links", filtersoundtracks.Links)
	page.Get(app, "/editor/soundtracks/lyrics/missing", filtersoundtracks.MissingLyrics)
	page.Get(app, "/editor/soundtracks/lyrics/unaligned", filtersoundtracks.UnalignedLyrics)
	page.Get(app, "/editor/soundtracks/tags", filtersoundtracks.Tags)
	page.Get(app, "/editor/soundtracks/file", filtersoundtracks.File)

	// Editor - Jobs
	page.Get(app, "/editor/jobs", jobs.Overview)

	// Log
	page.Get(app, "/log", editlog.Get)
	page.Get(app, "/log/from/:index", editlog.Get)
	page.Get(app, "/user/:nick/log", editlog.Get)
	page.Get(app, "/user/:nick/log/from/:index", editlog.Get)

	// Admin
	page.Get(app, "/admin", admin.Get)
	page.Get(app, "/admin/webdev", admin.WebDev)
	page.Get(app, "/admin/registrations", admin.UserRegistrations)
	page.Get(app, "/admin/errors/client", admin.ClientErrors)
	page.Get(app, "/admin/purchases", admin.PurchaseHistory)
	page.Get(app, "/admin/payments", admin.PaymentHistory)
}
