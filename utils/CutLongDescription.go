package utils

const maxDescriptionLength = 170

// CutLongDescription cuts a long description for use in OpenGraph tags.
func CutLongDescription(description string) string {
	if len(description) > maxDescriptionLength {
		return description[:maxDescriptionLength-3] + "..."
	}

	return description
}
