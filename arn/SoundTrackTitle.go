package arn

// SoundTrackTitle represents a song title.
type SoundTrackTitle struct {
	Canonical string `json:"canonical" editable:"true"`
	Native    string `json:"native" editable:"true"`
}

// String is the default representation of the title.
func (title *SoundTrackTitle) String() string {
	return title.ByUser(nil)
}

// ByUser returns the preferred title for the given user.
func (title *SoundTrackTitle) ByUser(user *User) string {
	if user == nil {
		if title.Canonical != "" {
			return title.Canonical
		}

		return title.Native
	}

	switch user.Settings().TitleLanguage {
	case "japanese":
		if title.Native == "" {
			return title.Canonical
		}

		return title.Native

	default:
		return title.ByUser(nil)
	}
}
