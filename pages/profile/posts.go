package profile

// import (
// 	"net/http"

// 	"github.com/aerogo/aero"
// 	"github.com/animenotifier/notify.moe/arn"
// 	"github.com/animenotifier/notify.moe/components"
// 	"github.com/animenotifier/notify.moe/utils"
// )

// const postLimit = 10

// // GetPostsByUser shows all forum posts of a particular user.
// func GetPostsByUser(ctx aero.Context) error {
// 	nick := ctx.Get("nick")
// 	viewUser, err := arn.GetUserByNick(nick)

// 	if err != nil {
// 		return ctx.Error(http.StatusNotFound, "User not found", err)
// 	}

// 	posts := viewUser.Posts()
// 	arn.SortPostsLatestFirst(posts)

// 	if len(posts) >= postLimit {
// 		posts = posts[:postLimit]
// 	}

// 	return ctx.HTML(components.LatestPosts(arn.ToPostables(posts), viewUser, arn.GetUserFromContext(ctx), ctx.Path()))

// }
