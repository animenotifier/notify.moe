package arn

const (
	// DefaultRating is the default rating value.
	DefaultRating = 0.0

	// AverageRating is the center rating in the system.
	// Note that the mathematically correct center would be a little higher,
	// but we don't care about these slight offsets.
	AverageRating = 5.0

	// MaxRating is the maximum rating users can give.
	MaxRating = 10.0

	// RatingCountThreshold is the number of users threshold that, when passed, doesn't dampen the result.
	RatingCountThreshold = 4
)

// AnimeRating represents the rating information for an anime.
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
