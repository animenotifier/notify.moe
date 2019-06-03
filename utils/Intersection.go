package utils

import "github.com/animenotifier/notify.moe/arn"

// Intersection returns common elements of a and b.
func Intersection(a []string, b []string) []string {
	var c []string

	for _, obj := range a {
		if arn.Contains(b, obj) {
			c = append(c, obj)
		}
	}

	return c
}
