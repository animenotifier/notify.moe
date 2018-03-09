package animediff

// JapaneseTitle describes differing Japanese titles.
type JapaneseTitle struct {
	TitleA string
	TitleB string
}

// Type returns the diff type.
func (diff *JapaneseTitle) Type() string {
	return "JapaneseTitle"
}

// Explanation returns the description.
func (diff *JapaneseTitle) Explanation() string {
	return "Japanese titles are different"
}

// DetailsA shows the details for the first anime.
func (diff *JapaneseTitle) DetailsA() string {
	return diff.TitleA
}

// DetailsB shows the details for the second anime.
func (diff *JapaneseTitle) DetailsB() string {
	return diff.TitleB
}
