package animediff

// CanonicalTitle describes differing titles.
type CanonicalTitle struct {
	TitleA string
	TitleB string
}

// Type returns the diff type.
func (diff *CanonicalTitle) Type() string {
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
