package arn

import (
	"encoding/json"

	"github.com/animenotifier/mal"
)

// MyAnimeListMatch ...
type MyAnimeListMatch struct {
	MyAnimeListItem *mal.AnimeListItem `json:"malItem"`
	ARNAnime        *Anime             `json:"arnAnime"`
}

// JSON ...
func (match *MyAnimeListMatch) JSON() string {
	b, err := json.Marshal(match)
	PanicOnError(err)
	return string(b)
}
