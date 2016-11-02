package main

import "sort"

// Genres ...
var Genres []string

// GenreIcons ...
var GenreIcons = map[string]string{
	"Action":        "bomb",
	"Adventure":     "diamond",
	"Cars":          "car",
	"Comedy":        "smile-o",
	"Drama":         "heartbeat",
	"Ecchi":         "heart-o",
	"Fantasy":       "tree",
	"Game":          "gamepad",
	"Harem":         "group",
	"Hentai":        "venus-mars",
	"Historical":    "history",
	"Horror":        "frown-o",
	"Kids":          "child",
	"Martial Arts":  "hand-rock-o",
	"Magic":         "magic",
	"Mecha":         "reddit-alien",
	"Military":      "fighter-jet",
	"Music":         "music",
	"Mystery":       "question",
	"Psychological": "lightbulb-o",
	"Romance":       "heart",
	"Sci-Fi":        "space-shuttle",
	"School":        "graduation-cap",
	"Seinen":        "male",
	"Shounen":       "male",
	"Shoujo":        "female",
	"Slice of Life": "hand-peace-o",
	"Sports":        "soccer-ball-o",
	"Supernatural":  "magic",
	"Super Power":   "flash",
	"Thriller":      "hourglass-end",
	"Vampire":       "eye",
}

func init() {
	for k := range GenreIcons {
		Genres = append(Genres, k)
	}
	sort.Strings(Genres)
}
