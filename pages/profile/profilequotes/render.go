package profilequotes

// import (
// 	"net/http"

// 	"github.com/aerogo/aero"
// 	"github.com/animenotifier/notify.moe/arn"
// 	"github.com/animenotifier/notify.moe/components"
// 	"github.com/animenotifier/notify.moe/utils"
// 	"github.com/animenotifier/notify.moe/utils/infinitescroll"
// )

// const (
// 	quotesFirstLoad = 12
// 	quotesPerScroll = 9
// )

// // render renders the quotes on user profiles.
// func render(ctx aero.Context, fetch func(userID string) []*arn.Quote) string {
// 	nick := ctx.Get("nick")
// 	index, _ := ctx.GetInt("index")
// 	user := arn.GetUserFromContext(ctx)
// 	viewUser, err := arn.GetUserByNick(nick)

// 	if err != nil {
// 		return ctx.Error(http.StatusNotFound, "User not found", err)
// 	}

// 	// Fetch all eligible quotes
// 	allQuotes := fetch(viewUser.ID)

// 	// Sort the quotes by publication date
// 	arn.SortQuotesLatestFirst(allQuotes)

// 	// Slice the part that we need
// 	quotes := allQuotes[index:]
// 	maxLength := quotesFirstLoad

// 	if index > 0 {
// 		maxLength = quotesPerScroll
// 	}

// 	if len(quotes) > maxLength {
// 		quotes = quotes[:maxLength]
// 	}

// 	// Next index
// 	nextIndex := infinitescroll.NextIndex(ctx, len(allQuotes), maxLength, index)

// 	// In case we're scrolling, send quotes only (without the page frame)
// 	if index > 0 {
// 		return ctx.HTML(components.QuotesScrollable(quotes, user))
// 	}

// 	// Otherwise, send the full page
// 	return ctx.HTML(components.ProfileQuotes(quotes, viewUser, nextIndex, user, ctx.Path()))
// }
