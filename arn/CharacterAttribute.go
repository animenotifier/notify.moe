package arn

// CharacterAttribute describes one attribute of a character, e.g. height or age.
type CharacterAttribute struct {
	Name  string `json:"name" editable:"true"`
	Value string `json:"value" editable:"true"`
}
