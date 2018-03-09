package utils

// AnimeTitleDiff describes differing titles.
type AnimeTitleDiff struct {
	TitleA string
	TitleB string
}

// String returns the description.
func (diff *AnimeTitleDiff) String() string {
	return "Titles are different"
}

// DetailsA shows the details for the first anime.
func (diff *AnimeTitleDiff) DetailsA() string {
	return diff.TitleA
}

// DetailsB shows the details for the second anime.
func (diff *AnimeTitleDiff) DetailsB() string {
	return diff.TitleB
}
