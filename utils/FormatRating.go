package utils

import "fmt"

// FormatRating formats the rating number.
func FormatRating(rating float64) string {
	return fmt.Sprintf("%.1f", rating)
}
