package profile

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// GetThreadsByUser shows all forum threads of a particular user.
func GetThreadsByUser(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	user, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	threads := user.Threads()
	arn.SortThreadsLatestFirst(threads)

	return ctx.HTML(components.ThreadList(threads))
}
