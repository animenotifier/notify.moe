package search

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/arn/search"
	"github.com/animenotifier/notify.moe/components"
)

const (
	maxUsers       = 25
	maxAnime       = 25
	maxPosts       = 2
	maxThreads     = 2
	maxSoundTracks = 3
	maxAMVs        = 3
	maxCharacters  = 22
	maxCompanies   = 3
)

// Get search page.
func Get(ctx aero.Context) error {
	term := ctx.Get("term")
	user := arn.GetUserFromContext(ctx)

	if len(term) > search.MaxSearchTermLength {
		ctx.SetStatus(http.StatusRequestEntityTooLarge)
	}

	users, animes, posts, threads, tracks, characters, amvs, companies := search.All(
		term,
		maxUsers,
		maxAnime,
		maxPosts,
		maxThreads,
		maxSoundTracks,
		maxCharacters,
		maxAMVs,
		maxCompanies,
	)

	return ctx.HTML(components.SearchResults(term, users, animes, posts, threads, tracks, characters, amvs, companies, nil, user))
}

// GetEmptySearch renders the search page with no contents.
func GetEmptySearch(ctx aero.Context) error {
	return ctx.HTML(components.SearchResults("", nil, nil, nil, nil, nil, nil, nil, nil, nil, arn.GetUserFromContext(ctx)))
}

// Anime search.
func Anime(ctx aero.Context) error {
	term := ctx.Get("term")
	user := arn.GetUserFromContext(ctx)

	if len(term) > search.MaxSearchTermLength {
		ctx.SetStatus(http.StatusRequestEntityTooLarge)
	}

	animes := search.Anime(term, maxAnime)
	return ctx.HTML(components.AnimeSearchResults(animes, user))
}

// Characters search.
func Characters(ctx aero.Context) error {
	term := ctx.Get("term")
	user := arn.GetUserFromContext(ctx)

	if len(term) > search.MaxSearchTermLength {
		ctx.SetStatus(http.StatusRequestEntityTooLarge)
	}

	characters := search.Characters(term, maxCharacters)
	return ctx.HTML(components.CharacterSearchResults(characters, user))
}

// Posts search.
func Posts(ctx aero.Context) error {
	term := ctx.Get("term")
	user := arn.GetUserFromContext(ctx)

	if len(term) > search.MaxSearchTermLength {
		ctx.SetStatus(http.StatusRequestEntityTooLarge)
	}

	posts := search.Posts(term, maxPosts)
	return ctx.HTML(components.PostsSearchResults(posts, user))
}

// Threads search.
func Threads(ctx aero.Context) error {
	term := ctx.Get("term")
	user := arn.GetUserFromContext(ctx)

	if len(term) > search.MaxSearchTermLength {
		ctx.SetStatus(http.StatusRequestEntityTooLarge)
	}

	threads := search.Threads(term, maxThreads)
	return ctx.HTML(components.ThreadsSearchResults(threads, user))
}

// SoundTracks search.
func SoundTracks(ctx aero.Context) error {
	term := ctx.Get("term")
	user := arn.GetUserFromContext(ctx)

	if len(term) > search.MaxSearchTermLength {
		ctx.SetStatus(http.StatusRequestEntityTooLarge)
	}

	tracks := search.SoundTracks(term, maxSoundTracks)
	return ctx.HTML(components.SoundTrackSearchResults(tracks, user))
}

// AMVs search.
func AMVs(ctx aero.Context) error {
	term := ctx.Get("term")
	user := arn.GetUserFromContext(ctx)

	if len(term) > search.MaxSearchTermLength {
		ctx.SetStatus(http.StatusRequestEntityTooLarge)
	}

	amvs := search.AMVs(term, maxAMVs)
	return ctx.HTML(components.AMVSearchResults(amvs, user))
}

// Users search.
func Users(ctx aero.Context) error {
	term := ctx.Get("term")

	if len(term) > search.MaxSearchTermLength {
		ctx.SetStatus(http.StatusRequestEntityTooLarge)
	}

	users := search.Users(term, maxUsers)
	return ctx.HTML(components.UserSearchResults(users))
}

// Companies search.
func Companies(ctx aero.Context) error {
	term := ctx.Get("term")

	if len(term) > search.MaxSearchTermLength {
		ctx.SetStatus(http.StatusRequestEntityTooLarge)
	}

	companies := search.Companies(term, maxCompanies)
	return ctx.HTML(components.CompanySearchResults(companies))
}
