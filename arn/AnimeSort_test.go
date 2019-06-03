package arn_test

import (
	"testing"

	"github.com/animenotifier/notify.moe/arn"
	"github.com/stretchr/testify/assert"
)

func TestAnimeSort(t *testing.T) {
	anime2011 := arn.FilterAnime(func(anime *arn.Anime) bool {
		return anime.StartDateTime().Year() == 2011
	})

	arn.SortAnimeByQuality(anime2011)

	// Best anime of 2011 needs to be Steins;Gate
	assert.Equal(t, "0KUWpFmig", anime2011[0].ID)
}
