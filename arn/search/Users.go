package search

import (
	"sort"
	"strings"

	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/arn/stringutils"
)

// Users searches all users.
func Users(originalTerm string, maxLength int) []*arn.User {
	term := strings.ToLower(stringutils.RemoveSpecialCharacters(originalTerm))
	results := make([]*Result, 0, maxLength)

	for user := range arn.StreamUsers() {
		if user.ID == originalTerm {
			return []*arn.User{user}
		}

		text := strings.ToLower(user.Nick)

		// Similarity check
		similarity := stringutils.AdvancedStringSimilarity(term, text)

		if similarity < MinStringSimilarity {
			continue
		}

		results = append(results, &Result{
			obj:        user,
			similarity: similarity,
		})
	}

	// Sort
	sort.Slice(results, func(i, j int) bool {
		return results[i].similarity > results[j].similarity
	})

	// Limit
	if len(results) >= maxLength {
		results = results[:maxLength]
	}

	// Final list
	final := make([]*arn.User, len(results))

	for i, result := range results {
		final[i] = result.obj.(*arn.User)
	}

	return final
}
