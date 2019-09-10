package arn

// Register a list of supported character roles.
func init() {
	DataLists["anime-character-roles"] = []*Option{
		{"main", "Main character"},
		{"supporting", "Supporting character"},
	}
}

// AnimeCharacter contains the information for a character and his role in an anime.
type AnimeCharacter struct {
	CharacterID string `json:"characterId" editable:"true"`
	Role        string `json:"role" editable:"true" datalist:"anime-character-roles"`
}

// Character returns the referenced character.
func (char *AnimeCharacter) Character() *Character {
	character, _ := GetCharacter(char.CharacterID)
	return character
}
