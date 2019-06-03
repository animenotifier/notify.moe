package characters

import "github.com/animenotifier/notify.moe/arn"

func fetchAll() []*arn.Character {
	return arn.FilterCharacters(func(character *arn.Character) bool {
		return !character.IsDraft
	})
}
