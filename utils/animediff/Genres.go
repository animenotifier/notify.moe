package animediff

import "strings"

// Genres describes differing genres.
type Genres struct {
	GenresA     []string
	GenresB     []string
	NumericHash uint64
}

// TypeName returns the diff type.
func (diff *Genres) TypeName() string {
	return "Genres"
}

// Explanation returns the description.
func (diff *Genres) Explanation() string {
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

// Hash returns the hash for the suggested value (from anime B).
func (diff *Genres) Hash() uint64 {
	return diff.NumericHash
}
