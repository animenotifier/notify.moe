package animediff

import "strings"

// Genres describes differing genres.
type Genres struct {
	GenresA []string
	GenresB []string
}

// String returns the description.
func (diff *Genres) String() string {
	return "Genres are different"
}

// DetailsA shows the details for the first anime.
func (diff *Genres) DetailsA() string {
	return strings.Join(diff.GenresA, ", ")
}

// DetailsB shows the details for the second anime.
func (diff *Genres) DetailsB() string {
	return strings.Join(diff.GenresB, ", ")
}
