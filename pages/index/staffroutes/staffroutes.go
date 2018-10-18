package staffroutes

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/layout"
	"github.com/animenotifier/notify.moe/pages/admin"
	"github.com/animenotifier/notify.moe/pages/editlog"
	"github.com/animenotifier/notify.moe/pages/editor"
	"github.com/animenotifier/notify.moe/pages/editor/filteranime"
	"github.com/animenotifier/notify.moe/pages/editor/filtercompanies"
	"github.com/animenotifier/notify.moe/pages/editor/filtersoundtracks"
	"github.com/animenotifier/notify.moe/pages/editor/jobs"
)

// Register registers the page routes.
func Register(l *layout.Layout) {
	// Editor
	l.Page("/editor", editor.Get)

	// Editor links can be filtered by year, status and type
	editorFilterable := func(route string, handler func(ctx *aero.Context) string) {
		l.Page(route+"/:year/:season/:status/:type", handler)
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
	l.Page("/editor/kitsu/new/anime", editor.NewKitsuAnime)

	// Editor - Companies
	l.Page("/editor/companies/description", filtercompanies.NoDescription)

	// Editor - Soundtracks
	l.Page("/editor/soundtracks/links", filtersoundtracks.Links)
	l.Page("/editor/soundtracks/lyrics/missing", filtersoundtracks.MissingLyrics)
	l.Page("/editor/soundtracks/lyrics/unaligned", filtersoundtracks.UnalignedLyrics)
	l.Page("/editor/soundtracks/tags", filtersoundtracks.Tags)
	l.Page("/editor/soundtracks/file", filtersoundtracks.File)

	// Editor - Jobs
	l.Page("/editor/jobs", jobs.Overview)

	// Log
	l.Page("/log", editlog.Get)
	l.Page("/log/from/:index", editlog.Get)
	l.Page("/user/:nick/log", editlog.Get)
	l.Page("/user/:nick/log/from/:index", editlog.Get)

	// Admin
	l.Page("/admin", admin.Get)
	l.Page("/admin/webdev", admin.WebDev)
	l.Page("/admin/registrations", admin.UserRegistrations)
	l.Page("/admin/errors/client", admin.ClientErrors)
	l.Page("/admin/purchases", admin.PurchaseHistory)
	l.Page("/admin/payments", admin.PaymentHistory)
}
