package arn

import (
	"github.com/animenotifier/mal"
	jsoniter "github.com/json-iterator/go"
)

// MyAnimeListMatch ...
type MyAnimeListMatch struct {
	MyAnimeListItem *mal.AnimeListItem `json:"malItem"`
	ARNAnime        *Anime             `json:"arnAnime"`
}

// JSON ...
func (match *MyAnimeListMatch) JSON() string {
	b, err := jsoniter.Marshal(match)
	PanicOnError(err)
	return string(b)
}
