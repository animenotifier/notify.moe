package animelistitem

import (
	"errors"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get anime page.
func Get(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	nick := ctx.Get("nick")
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	animeList := viewUser.AnimeList()

	if animeList == nil {
		return ctx.Error(http.StatusNotFound, "Anime list not found", err)
	}

	animeID := ctx.Get("id")
	item := animeList.Find(animeID)

	if item == nil {
		return ctx.Error(http.StatusNotFound, "List item not found", errors.New("This anime does not exist in "+viewUser.Nick+"'s anime list"))
	}

	anime := item.Anime()

	return ctx.HTML(components.AnimeListItem(animeList.User(), item, anime, user))
}

// t := reflect.TypeOf(item).Elem()
// v := reflect.ValueOf(item).Elem()

// for i := 0; i < t.NumField(); i++ {
// 	fieldInfo := t.Field(i)

// 	if fieldInfo.Anonymous || unicode.IsLower([]rune(fieldInfo.Name)[0]) {
// 		continue
// 	}

// 	fmt.Println(fieldInfo.Name, v.Field(i).Interface())
// }
