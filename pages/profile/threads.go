package profile

// import (
// 	"net/http"

// 	"github.com/aerogo/aero"
// 	"github.com/animenotifier/notify.moe/arn"
// 	"github.com/animenotifier/notify.moe/components"
// 	"github.com/animenotifier/notify.moe/utils"
// )

// const maxThreads = 20

// // GetThreadsByUser shows all forum threads of a particular user.
// func GetThreadsByUser(ctx aero.Context) error {
// 	nick := ctx.Get("nick")
// 	viewUser, err := arn.GetUserByNick(nick)

// 	if err != nil {
// 		return ctx.Error(http.StatusNotFound, "User not found", err)
// 	}

// 	threads := viewUser.Threads()
// 	arn.SortThreadsLatestFirst(threads)

// 	if len(threads) > maxThreads {
// 		threads = threads[:maxThreads]
// 	}

// 	return ctx.HTML(components.ProfileThreads(threads, viewUser, arn.GetUserFromContext(ctx), ctx.Path()))
// }
