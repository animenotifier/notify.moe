package search

import (
	"sort"
	"strings"

	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/arn/stringutils"
)

// Characters searches all characters.
func Characters(originalTerm string, maxLength int) []*arn.Character {
	if maxLength == 0 {
		return nil
	}

	term := strings.ToLower(stringutils.RemoveSpecialCharacters(originalTerm))
	termHasUnicode := stringutils.ContainsUnicodeLetters(term)
	results := make([]*Result, 0, maxLength)

	for character := range arn.StreamCharacters() {
		if character.ID == originalTerm {
			return []*arn.Character{character}
		}

		if character.Image.Extension == "" {
			continue
		}

		// Canonical
		text := strings.ToLower(stringutils.RemoveSpecialCharacters(character.Name.Canonical))

		if text == term {
			results = append(results, &Result{
				obj:        character,
				similarity: float64(20 + len(character.Likes)),
			})
			continue
		}

		spaceCount := 0
		start := 0
		found := false

		for i := 0; i <= len(text); i++ {
			if i == len(text) || text[i] == ' ' {
				part := text[start:i]

				if part == term {
					results = append(results, &Result{
						obj:        character,
						similarity: float64(10 - spaceCount*5 + len(character.Likes)),
					})

					found = true
					break
				}

				start = i + 1
				spaceCount++
			}
		}

		if found {
			continue
		}

		// Japanese
		if termHasUnicode {
			if strings.Contains(character.Name.Japanese, term) {
				results = append(results, &Result{
					obj:        character,
					similarity: float64(len(character.Likes)),
				})
				continue
			}
		}
	}

	// Sort
	sort.Slice(results, func(i, j int) bool {
		similarityA := results[i].similarity
		similarityB := results[j].similarity

		if similarityA == similarityB {
			characterA := results[i].obj.(*arn.Character)
			characterB := results[j].obj.(*arn.Character)

			if characterA.Name.Canonical == characterB.Name.Canonical {
				return characterA.ID < characterB.ID
			}

			return characterA.Name.Canonical < characterB.Name.Canonical
		}

		return similarityA > similarityB
	})

	// Limit
	if len(results) >= maxLength {
		results = results[:maxLength]
	}

	// Final list
	final := make([]*arn.Character, len(results))

	for i, result := range results {
		final[i] = result.obj.(*arn.Character)
	}

	return final
}
