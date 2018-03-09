package animediff

// ShorterSynopsis describes differing Japanese titles.
type ShorterSynopsis struct {
	SynopsisA string
	SynopsisB string
}

// String returns the description.
func (diff *ShorterSynopsis) String() string {
	return "Synopsis is shorter"
}

// DetailsA shows the details for the first anime.
func (diff *ShorterSynopsis) DetailsA() string {
	return diff.SynopsisA
}

// DetailsB shows the details for the second anime.
func (diff *ShorterSynopsis) DetailsB() string {
	return diff.SynopsisB
}
