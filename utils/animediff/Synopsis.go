package animediff

// Synopsis describes differing synopsis.
type Synopsis struct {
	SynopsisA   string
	SynopsisB   string
	NumericHash uint64
}

// TypeName returns the diff type.
func (diff *Synopsis) TypeName() string {
	return "Synopsis"
}

// Explanation returns the description.
func (diff *Synopsis) Explanation() string {
	return "Synopsis is shorter"
}

// DetailsA shows the details for the first anime.
func (diff *Synopsis) DetailsA() string {
	return diff.SynopsisA
}

// DetailsB shows the details for the second anime.
func (diff *Synopsis) DetailsB() string {
	return diff.SynopsisB
}

// Hash returns the hash for the suggested value (from anime B).
func (diff *Synopsis) Hash() uint64 {
	return diff.NumericHash
}
