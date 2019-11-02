package arn

import (
	"encoding/json"

	"github.com/animenotifier/kitsu"
)

// KitsuMatch ...
type KitsuMatch struct {
	KitsuItem *kitsu.LibraryEntry `json:"kitsuItem"`
	ARNAnime  *Anime              `json:"arnAnime"`
}

// JSON ...
func (match *KitsuMatch) JSON() string {
	b, err := json.Marshal(match)
	PanicOnError(err)
	return string(b)
}
