package arn

// CharacterFinder holds an internal map of ID to anime mappings
// and is therefore very efficient to use when trying to find
// anime by a given service and ID.
type CharacterFinder struct {
	idToCharacter map[string]*Character
	mappingName   string
}

// NewCharacterFinder creates a new finder for external characters.
func NewCharacterFinder(mappingName string) *CharacterFinder {
	finder := &CharacterFinder{
		idToCharacter: map[string]*Character{},
		mappingName:   mappingName,
	}

	for character := range StreamCharacters() {
		finder.Add(character)
	}

	return finder
}

// Add adds a character to the search pool.
func (finder *CharacterFinder) Add(character *Character) {
	id := character.GetMapping(finder.mappingName)

	if id != "" {
		finder.idToCharacter[id] = character
	}
}

// GetCharacter tries to find an external anime in our anime database.
func (finder *CharacterFinder) GetCharacter(id string) *Character {
	return finder.idToCharacter[id]
}
