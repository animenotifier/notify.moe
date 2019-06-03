package arn

// AnimeTitle ...
type AnimeTitle struct {
	Canonical string   `json:"canonical" editable:"true"`
	Romaji    string   `json:"romaji" editable:"true"`
	English   string   `json:"english" editable:"true"`
	Japanese  string   `json:"japanese" editable:"true"`
	Hiragana  string   `json:"hiragana" editable:"true"`
	Synonyms  []string `json:"synonyms" editable:"true"`
}

// ByUser returns the preferred title for the given user.
func (title *AnimeTitle) ByUser(user *User) string {
	if user == nil {
		return title.Canonical
	}

	switch user.Settings().TitleLanguage {
	case "canonical":
		return title.Canonical
	case "romaji":
		if title.Romaji == "" {
			return title.Canonical
		}

		return title.Romaji
	case "english":
		if title.English == "" {
			return title.Canonical
		}

		return title.English
	case "japanese":
		if title.Japanese == "" {
			return title.Canonical
		}

		return title.Japanese
	default:
		panic("Invalid title language")
	}
}
