package animediff

// CanonicalTitle describes differing titles.
type CanonicalTitle struct {
	TitleA string
	TitleB string
}

// String returns the description.
func (diff *CanonicalTitle) String() string {
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
