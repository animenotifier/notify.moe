package utils

import "strings"

// AnimeGenresDiff describes differing genres.
type AnimeGenresDiff struct {
	GenresA []string
	GenresB []string
}

// String returns the description.
func (diff *AnimeGenresDiff) String() string {
	return "Genres are different"
}

// DetailsA shows the details for the first anime.
func (diff *AnimeGenresDiff) DetailsA() string {
	return strings.Join(diff.GenresA, ", ")
}

// DetailsB shows the details for the second anime.
func (diff *AnimeGenresDiff) DetailsB() string {
	return strings.Join(diff.GenresB, ", ")
}
