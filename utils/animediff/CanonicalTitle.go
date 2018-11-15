package animediff

// CanonicalTitle describes differing titles.
type CanonicalTitle struct {
	TitleA      string
	TitleB      string
	NumericHash uint64
}

// TypeName returns the diff type.
func (diff *CanonicalTitle) TypeName() string {
	return "CanonicalTitle"
}

// Explanation returns the description.
func (diff *CanonicalTitle) Explanation() string {
	return "Canonical titles are different"
}

// DetailsA shows the details for the first anime.
func (diff *CanonicalTitle) DetailsA() string {
	return diff.TitleA
}

// DetailsB shows the details for the second anime.
func (diff *CanonicalTitle) DetailsB() string {
	return diff.TitleB
}

// Hash returns the hash for the suggested value (from anime B).
func (diff *CanonicalTitle) Hash() uint64 {
	return diff.NumericHash
}
