package characters

import "github.com/animenotifier/arn"

func fetchAll() []*arn.Character {
	return arn.FilterCharacters(func(character *arn.Character) bool {
		return !character.IsDraft
	})
}
