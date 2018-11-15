package animediff

// RomajiTitle describes differing Romaji titles.
type RomajiTitle struct {
	TitleA      string
	TitleB      string
	NumericHash uint64
}

// TypeName returns the diff type.
func (diff *RomajiTitle) TypeName() string {
	return "RomajiTitle"
}

// Explanation returns the description.
func (diff *RomajiTitle) Explanation() string {
	return "Romaji titles are different"
}

// DetailsA shows the details for the first anime.
func (diff *RomajiTitle) DetailsA() string {
	return diff.TitleA
}

// DetailsB shows the details for the second anime.
func (diff *RomajiTitle) DetailsB() string {
	return diff.TitleB
}

// Hash returns the hash for the suggested value (from anime B).
func (diff *RomajiTitle) Hash() uint64 {
	return diff.NumericHash
}
