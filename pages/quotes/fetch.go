package quotes

import "github.com/animenotifier/notify.moe/arn"

// fetchAll returns all quotes
func fetchAll() []*arn.Quote {
	return arn.FilterQuotes(func(quote *arn.Quote) bool {
		return !quote.IsDraft && len(quote.Text.English) > 0
	})
}
