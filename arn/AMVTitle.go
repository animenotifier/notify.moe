package arn

// AMVTitle is the same as a soundtrack title.
type AMVTitle SoundTrackTitle

// String is the default representation of the title.
func (title *AMVTitle) String() string {
	return title.ByUser(nil)
}

// ByUser returns the preferred title for the given user.
func (title *AMVTitle) ByUser(user *User) string {
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
