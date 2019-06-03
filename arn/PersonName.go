package arn

// PersonName represents the name of a person.
type PersonName struct {
	English  Name `json:"english" editable:"true"`
	Japanese Name `json:"japanese" editable:"true"`
}

// String returns the default visualization of the name.
func (name *PersonName) String() string {
	return name.ByUser(nil)
}

// ByUser returns the preferred name for the given user.
func (name *PersonName) ByUser(user *User) string {
	if user == nil {
		return name.English.String()
	}

	switch user.Settings().TitleLanguage {
	case "japanese":
		if name.Japanese.String() == "" {
			return name.English.String()
		}

		return name.Japanese.String()

	default:
		return name.English.String()
	}
}
