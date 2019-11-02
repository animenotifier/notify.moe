package arn

import (
	"encoding/json"

	"github.com/animenotifier/anilist"
)

// AniListMatch ...
type AniListMatch struct {
	AniListItem *anilist.AnimeListItem `json:"anilistItem"`
	ARNAnime    *Anime                 `json:"arnAnime"`
}

// JSON ...
func (match *AniListMatch) JSON() string {
	b, err := json.Marshal(match)
	PanicOnError(err)
	return string(b)
}
