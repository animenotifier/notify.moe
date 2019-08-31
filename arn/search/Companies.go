package search

import (
	"sort"
	"strings"

	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/arn/stringutils"
)

// Companies searches all companies.
func Companies(originalTerm string, maxLength int) []*arn.Company {
	term := strings.ToLower(stringutils.RemoveSpecialCharacters(originalTerm))
	results := make([]*Result, 0, maxLength)

	for company := range arn.StreamCompanies() {
		if company.ID == originalTerm {
			return []*arn.Company{company}
		}

		if company.IsDraft {
			continue
		}

		text := strings.ToLower(stringutils.RemoveSpecialCharacters(company.Name.English))
		similarity := stringutils.AdvancedStringSimilarity(term, text)

		if similarity >= MinStringSimilarity {
			results = append(results, &Result{
				obj:        company,
				similarity: similarity,
			})
		}
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
	final := make([]*arn.Company, len(results))

	for i, result := range results {
		final[i] = result.obj.(*arn.Company)
	}

	return final
}
