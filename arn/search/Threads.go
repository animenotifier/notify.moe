package search

import (
	"sort"
	"strings"

	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/arn/stringutils"
)

// Threads searches all threads.
func Threads(originalTerm string, maxLength int) []*arn.Thread {
	term := strings.ToLower(stringutils.RemoveSpecialCharacters(originalTerm))
	results := make([]*arn.Thread, 0, maxLength)

	for thread := range arn.StreamThreads() {
		if thread.ID == originalTerm {
			return []*arn.Thread{thread}
		}

		text := strings.ToLower(thread.Text)

		if strings.Contains(text, term) {
			results = append(results, thread)
			continue
		}

		text = strings.ToLower(thread.Title)

		if strings.Contains(text, term) {
			results = append(results, thread)
			continue
		}
	}

	// Sort
	sort.Slice(results, func(i, j int) bool {
		return results[i].Created > results[j].Created
	})

	// Limit
	if len(results) >= maxLength {
		results = results[:maxLength]
	}

	return results
}
