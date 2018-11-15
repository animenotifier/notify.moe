package animediff

// Status describes differing Romaji titles.
type Status struct {
	StatusA     string
	StatusB     string
	NumericHash uint64
}

// TypeName returns the diff type.
func (diff *Status) TypeName() string {
	return "Status"
}

// Explanation returns the description.
func (diff *Status) Explanation() string {
	return "Status is different"
}

// DetailsA shows the details for the first anime.
func (diff *Status) DetailsA() string {
	return diff.StatusA
}

// DetailsB shows the details for the second anime.
func (diff *Status) DetailsB() string {
	return diff.StatusB
}

// Hash returns the hash for the suggested value (from anime B).
func (diff *Status) Hash() uint64 {
	return diff.NumericHash
}
