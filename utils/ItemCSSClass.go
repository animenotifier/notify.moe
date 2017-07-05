package utils

import (
	"github.com/animenotifier/arn"
)

// ItemCSSClass removes mountable class if the list has too many items.
func ItemCSSClass(list *arn.AnimeList, index int) string {
	if index > 20 || len(list.Items) > 50 {
		return "anime-list-item"
	}

	return "anime-list-item mountable"
}
