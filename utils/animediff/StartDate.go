package animediff

// StartDate describes differing Romaji titles.
type StartDate struct {
	DateA       string
	DateB       string
	NumericHash uint64
}

// TypeName returns the diff type.
func (diff *StartDate) TypeName() string {
	return "StartDate"
}

// Explanation returns the description.
func (diff *StartDate) Explanation() string {
	return "Start dates are different"
}

// DetailsA shows the details for the first anime.
func (diff *StartDate) DetailsA() string {
	return diff.DateA
}

// DetailsB shows the details for the second anime.
func (diff *StartDate) DetailsB() string {
	return diff.DateB
}

// Hash returns the hash for the suggested value (from anime B).
func (diff *StartDate) Hash() uint64 {
	return diff.NumericHash
}
