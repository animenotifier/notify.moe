package search

import (
	"strings"

	"github.com/animenotifier/notify.moe/utils"

	"github.com/aerogo/flow"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/arn/search"
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
	user := utils.GetUser(ctx)

	users, animes, posts, threads, tracks, characters, companies := search.All(
		term,
		maxUsers,
		maxAnime,
		maxPosts,
		maxThreads,
		maxSoundTracks,
		maxCharacters,
		maxCompanies,
	)

	return ctx.HTML(components.SearchResults(term, users, animes, posts, threads, tracks, characters, companies, nil, user))
}

// GetEmptySearch renders the search page with no contents.
func GetEmptySearch(ctx *aero.Context) string {
	return ctx.HTML(components.SearchResults("", nil, nil, nil, nil, nil, nil, nil, nil, utils.GetUser(ctx)))
}

// Anime search.
func Anime(ctx *aero.Context) string {
	term := ctx.Get("term")
	term = strings.TrimPrefix(term, "/")

	animes := search.Anime(term, maxAnime)
	return ctx.HTML(components.AnimeSearchResults(animes))
}

// Characters search.
func Characters(ctx *aero.Context) string {
	term := ctx.Get("term")
	term = strings.TrimPrefix(term, "/")
	user := utils.GetUser(ctx)

	characters := search.Characters(term, maxCharacters)
	return ctx.HTML(components.CharacterSearchResults(characters, user))
}

// Forum search.
func Forum(ctx *aero.Context) string {
	term := ctx.Get("term")
	term = strings.TrimPrefix(term, "/")

	var posts []*arn.Post
	var threads []*arn.Thread

	flow.Parallel(func() {
		posts = search.Posts(term, maxPosts)
	}, func() {
		threads = search.Threads(term, maxThreads)
	})

	return ctx.HTML(components.ForumSearchResults(posts, threads))
}

// SoundTracks search.
func SoundTracks(ctx *aero.Context) string {
	term := ctx.Get("term")
	term = strings.TrimPrefix(term, "/")

	tracks := search.SoundTracks(term, maxSoundTracks)
	return ctx.HTML(components.SoundTrackSearchResults(tracks))
}

// Users search.
func Users(ctx *aero.Context) string {
	term := ctx.Get("term")
	term = strings.TrimPrefix(term, "/")

	users := search.Users(term, maxUsers)
	return ctx.HTML(components.UserSearchResults(users))
}

// Companies search.
func Companies(ctx *aero.Context) string {
	term := ctx.Get("term")
	term = strings.TrimPrefix(term, "/")

	companies := search.Companies(term, maxCompanies)
	return ctx.HTML(components.CompanySearchResults(companies))
}
