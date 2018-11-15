package animediff

// EndDate describes differing Romaji titles.
type EndDate struct {
	DateA       string
	DateB       string
	NumericHash uint64
}

// TypeName returns the diff type.
func (diff *EndDate) TypeName() string {
	return "EndDate"
}

// Explanation returns the description.
func (diff *EndDate) Explanation() string {
	return "End dates are different"
}

// DetailsA shows the details for the first anime.
func (diff *EndDate) DetailsA() string {
	return diff.DateA
}

// DetailsB shows the details for the second anime.
func (diff *EndDate) DetailsB() string {
	return diff.DateB
}

// Hash returns the hash for the suggested value (from anime B).
func (diff *EndDate) Hash() uint64 {
	return diff.NumericHash
}
