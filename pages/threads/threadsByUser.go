package threads

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// GetByUser ...
func GetByUser(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	user, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	threads := user.Threads()
	arn.SortThreadsByDate(threads)

	return ctx.HTML(components.ThreadList(threads))
}
