package arn

import (
	"github.com/animenotifier/anilist"
	jsoniter "github.com/json-iterator/go"
)

// AniListMatch ...
type AniListMatch struct {
	AniListItem *anilist.AnimeListItem `json:"anilistItem"`
	ARNAnime    *Anime                 `json:"arnAnime"`
}

// JSON ...
func (match *AniListMatch) JSON() string {
	b, err := jsoniter.Marshal(match)
	PanicOnError(err)
	return string(b)
}
