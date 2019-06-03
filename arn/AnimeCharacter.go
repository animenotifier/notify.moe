package arn

// Register a list of supported character roles.
func init() {
	DataLists["anime-character-roles"] = []*Option{
		{"main", "Main character"},
		{"supporting", "Supporting character"},
	}
}

// AnimeCharacter ...
type AnimeCharacter struct {
	CharacterID string `json:"characterId" editable:"true"`
	Role        string `json:"role" editable:"true" datalist:"anime-character-roles"`
}

// Character ...
func (char *AnimeCharacter) Character() *Character {
	character, _ := GetCharacter(char.CharacterID)
	return character
}
