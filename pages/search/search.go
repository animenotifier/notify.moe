package search

import (
	"strings"

	"github.com/aerogo/flow"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const (
	maxUsers       = 25
	maxAnime       = 25
	maxPosts       = 2
	maxThreads     = 2
	maxSoundTracks = 3
	maxCharacters  = 22
	maxCompanies   = 3
)

// Get search page.
func Get(ctx *aero.Context) string {
	term := ctx.Get("term")
	term = strings.TrimPrefix(term, "/")

	users, animes, posts, threads, tracks, characters, companies := arn.Search(
		term,
		maxUsers,
		maxAnime,
		maxPosts,
		maxThreads,
		maxSoundTracks,
		maxCharacters,
		maxCompanies,
	)

	return ctx.HTML(components.SearchResults(term, users, animes, posts, threads, tracks, characters, companies, nil))
}

// GetEmptySearch renders the search page with no contents.
func GetEmptySearch(ctx *aero.Context) string {
	return ctx.HTML(components.SearchResults("", nil, nil, nil, nil, nil, nil, nil, nil))
}

// Anime search.
func Anime(ctx *aero.Context) string {
	term := ctx.Get("term")
	term = strings.TrimPrefix(term, "/")

	animes := arn.SearchAnime(term, maxAnime)
	return ctx.HTML(components.AnimeSearchResults(animes))
}

// Characters search.
func Characters(ctx *aero.Context) string {
	term := ctx.Get("term")
	term = strings.TrimPrefix(term, "/")

	characters := arn.SearchCharacters(term, maxCharacters)
	return ctx.HTML(components.CharacterSearchResults(characters))
}

// Forum search.
func Forum(ctx *aero.Context) string {
	term := ctx.Get("term")
	term = strings.TrimPrefix(term, "/")

	var posts []*arn.Post
	var threads []*arn.Thread

	flow.Parallel(func() {
		posts = arn.SearchPosts(term, maxPosts)
	}, func() {
		threads = arn.SearchThreads(term, maxThreads)
	})

	return ctx.HTML(components.ForumSearchResults(posts, threads))
}

// SoundTracks search.
func SoundTracks(ctx *aero.Context) string {
	term := ctx.Get("term")
	term = strings.TrimPrefix(term, "/")

	tracks := arn.SearchSoundTracks(term, maxSoundTracks)
	return ctx.HTML(components.SoundTrackSearchResults(tracks))
}

// Users search.
func Users(ctx *aero.Context) string {
	term := ctx.Get("term")
	term = strings.TrimPrefix(term, "/")

	users := arn.SearchUsers(term, maxUsers)
	return ctx.HTML(components.UserSearchResults(users))
}

// Companies search.
func Companies(ctx *aero.Context) string {
	term := ctx.Get("term")
	term = strings.TrimPrefix(term, "/")

	companies := arn.SearchCompanies(term, maxCompanies)
	return ctx.HTML(components.CompanySearchResults(companies))
}
