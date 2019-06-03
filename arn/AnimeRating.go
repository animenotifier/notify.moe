package arn

// DefaultRating is the default rating value.
const DefaultRating = 0.0

// AverageRating is the center rating in the system.
// Note that the mathematically correct center would be a little higher,
// but we don't care about these slight offsets.
const AverageRating = 5.0

// MaxRating is the maximum rating users can give.
const MaxRating = 10.0

// RatingCountThreshold is the number of users threshold that, when passed, doesn't dampen the result.
const RatingCountThreshold = 4

// AnimeRating ...
type AnimeRating struct {
	AnimeListItemRating

	// The amount of people who rated
	Count AnimeRatingCount `json:"count"`
}

// AnimeRatingCount ...
type AnimeRatingCount struct {
	Overall    int `json:"overall"`
	Story      int `json:"story"`
	Visuals    int `json:"visuals"`
	Soundtrack int `json:"soundtrack"`
}
