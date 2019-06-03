package arn

import (
	"github.com/animenotifier/kitsu"
	jsoniter "github.com/json-iterator/go"
)

// KitsuMatch ...
type KitsuMatch struct {
	KitsuItem *kitsu.LibraryEntry `json:"kitsuItem"`
	ARNAnime  *Anime              `json:"arnAnime"`
}

// JSON ...
func (match *KitsuMatch) JSON() string {
	b, err := jsoniter.Marshal(match)
	PanicOnError(err)
	return string(b)
}
