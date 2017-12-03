package quotes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get renders the quotes page.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	quotes := arn.AllQuotes()
	return ctx.HTML(components.Quotes(quotes, user))
}
