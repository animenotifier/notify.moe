package arn_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/animenotifier/notify.moe/arn"
)

func TestNewAnime(t *testing.T) {
	anime := arn.NewAnime()
	assert.NotNil(t, anime)
	assert.NotEqual(t, anime.ID, "")
	assert.NotEqual(t, anime.Created, "")
}

func TestGetAnime(t *testing.T) {
	// Existing anime
	anime, err := arn.GetAnime("74y2cFiiR")
	assert.Nil(t, err)
	assert.NotNil(t, anime)
	assert.NotEqual(t, anime.ID, "")
	assert.NotEqual(t, anime.Title.Canonical, "")

	// Not existing anime
	anime, err = arn.GetAnime("does not exist")
	assert.NotNil(t, err)
	assert.Nil(t, anime)
}

func TestAllAnime(t *testing.T) {
	validAnimeStatus := []string{
		"finished",
		"current",
		"upcoming",
		"tba",
	}

	validAnimeType := []string{
		"tv",
		"movie",
		"ova",
		"ona",
		"special",
		"music",
	}

	allAnime := arn.AllAnime()

	for _, anime := range allAnime {
		assert.NotEqual(t, anime.ID, "")
		assert.Contains(t, validAnimeStatus, anime.Status)
		assert.Contains(t, validAnimeType, anime.Type)
		assert.Contains(t, validAnimeStatus, anime.CalculatedStatus())
		assert.NotEqual(t, anime.StatusHumanReadable(), "")
		assert.NotEqual(t, anime.TypeHumanReadable(), "")
		assert.NotEqual(t, anime.Link(), "")
		assert.NotEqual(t, anime.EpisodeCountString(), "")

		anime.Episodes()
		anime.Characters()
		anime.StartDateTime()
		anime.EndDateTime()
		anime.HasImage()
		anime.GetMapping("shoboi/anime")
		anime.Studios()
		anime.Producers()
		anime.Licensors()
		anime.Prequels()
	}
}
