package search

import (
	"github.com/aerogo/flow"
	"github.com/animenotifier/notify.moe/arn"
)

// MinimumStringSimilarity is the minimum JaroWinkler distance we accept for search results.
const MinimumStringSimilarity = 0.89

// popularityDamping reduces the factor of popularity in search results.
const popularityDamping = 0.0009

// Result ...
type Result struct {
	obj        interface{}
	similarity float64
}

// All is a fuzzy search.
func All(term string, maxUsers, maxAnime, maxPosts, maxThreads, maxTracks, maxCharacters, maxAMVs, maxCompanies int) ([]*arn.User, []*arn.Anime, []*arn.Post, []*arn.Thread, []*arn.SoundTrack, []*arn.Character, []*arn.AMV, []*arn.Company) {
	if term == "" {
		return nil, nil, nil, nil, nil, nil, nil, nil
	}

	var (
		userResults      []*arn.User
		animeResults     []*arn.Anime
		postResults      []*arn.Post
		threadResults    []*arn.Thread
		trackResults     []*arn.SoundTrack
		characterResults []*arn.Character
		amvResults       []*arn.AMV
		companyResults   []*arn.Company
	)

	flow.Parallel(func() {
		userResults = Users(term, maxUsers)
	}, func() {
		animeResults = Anime(term, maxAnime)
	}, func() {
		postResults = Posts(term, maxPosts)
	}, func() {
		threadResults = Threads(term, maxThreads)
	}, func() {
		trackResults = SoundTracks(term, maxTracks)
	}, func() {
		characterResults = Characters(term, maxCharacters)
	}, func() {
		amvResults = AMVs(term, maxAMVs)
	}, func() {
		companyResults = Companies(term, maxCompanies)
	})

	return userResults, animeResults, postResults, threadResults, trackResults, characterResults, amvResults, companyResults
}
