package quotes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils/infinitescroll"
)

// Best renders the best quotes.
func Best(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	index, _ := ctx.GetInt("index")

	// Fetch all eligible quotes
	allQuotes := fetchAll()

	// Sort the quotes by number of likes
	arn.SortQuotesPopularFirst(allQuotes)

	// Slice the part that we need
	quotes := allQuotes[index:]

	if len(quotes) > maxQuotes {
		quotes = quotes[:maxQuotes]
	}

	// Next index
	nextIndex := infinitescroll.NextIndex(ctx, len(allQuotes), maxQuotes, index)

	// In case we're scrolling, send quotes only (without the page frame)
	if index > 0 {
		return ctx.HTML(components.QuotesScrollable(quotes, user))
	}

	// Otherwise, send the full page
	return ctx.HTML(components.Quotes(quotes, nextIndex, user))
}
