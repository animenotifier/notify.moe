package arn

// CharacterName ...
type CharacterName struct {
	Canonical string   `json:"canonical" editable:"true"`
	English   string   `json:"english" editable:"true"`
	Japanese  string   `json:"japanese" editable:"true"`
	Synonyms  []string `json:"synonyms" editable:"true"`
}

// ByUser returns the preferred name for the given user.
func (name *CharacterName) ByUser(user *User) string {
	if user == nil {
		return name.Canonical
	}

	switch user.Settings().TitleLanguage {
	case "canonical", "romaji":
		return name.Canonical
	case "english":
		if name.English == "" {
			return name.Canonical
		}

		return name.English
	case "japanese":
		if name.Japanese == "" {
			return name.Canonical
		}

		return name.Japanese
	default:
		panic("Invalid name language")
	}
}
