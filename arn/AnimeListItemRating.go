package arn

// AnimeListItemRating ...
type AnimeListItemRating struct {
	Overall    float64 `json:"overall" editable:"true"`
	Story      float64 `json:"story" editable:"true"`
	Visuals    float64 `json:"visuals" editable:"true"`
	Soundtrack float64 `json:"soundtrack" editable:"true"`
}

// IsNotRated tells you whether all ratings are zero.
func (rating *AnimeListItemRating) IsNotRated() bool {
	return rating.Overall == 0 && rating.Story == 0 && rating.Visuals == 0 && rating.Soundtrack == 0
}

// Reset sets all values to the default anime average rating.
func (rating *AnimeListItemRating) Reset() {
	rating.Overall = DefaultRating
	rating.Story = DefaultRating
	rating.Visuals = DefaultRating
	rating.Soundtrack = DefaultRating
}

// Clamp ...
func (rating *AnimeListItemRating) Clamp() {
	if rating.Overall < 0 {
		rating.Overall = 0
	}

	if rating.Story < 0 {
		rating.Story = 0
	}

	if rating.Visuals < 0 {
		rating.Visuals = 0
	}

	if rating.Soundtrack < 0 {
		rating.Soundtrack = 0
	}

	if rating.Overall > MaxRating {
		rating.Overall = MaxRating
	}

	if rating.Story > MaxRating {
		rating.Story = MaxRating
	}

	if rating.Visuals > MaxRating {
		rating.Visuals = MaxRating
	}

	if rating.Soundtrack > MaxRating {
		rating.Soundtrack = MaxRating
	}
}
