package arn

import "sort"

// Genres ...
var Genres []string

// Icons
var genreIcons = map[string]string{
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
	"Mecha":         "mecha",
	"Military":      "fighter-jet",
	"Music":         "music",
	"Mystery":       "question",
	"Psychological": "lightbulb-o",
	"Romance":       "heart",
	"Sci-Fi":        "rocket",
	"School":        "graduation-cap",
	"Seinen":        "male",
	"Shounen":       "child",
	"Shoujo":        "female",
	"Slice of Life": "hand-peace-o",
	"Space":         "space-shuttle",
	"Sports":        "soccer-ball-o",
	"Supernatural":  "magic",
	"Super Power":   "flash",
	"Thriller":      "hourglass-end",
	"Vampire":       "eye",
}

// GetGenreIcon returns the unprefixed icon class name for the genre.
func GetGenreIcon(genre string) string {
	icon, exists := genreIcons[genre]

	if exists {
		return icon
	}

	return "circle"
}

func init() {
	for k := range genreIcons {
		Genres = append(Genres, k)
	}
	sort.Strings(Genres)
}
