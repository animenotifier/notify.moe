package main

import (
	"github.com/aerogo/api"
	"github.com/animenotifier/arn"
)

func init() {
	api := api.New("/api/", arn.DB)
	api.Install(app)
}

// func init() {
// 	// app.Get("/all/anime", func(ctx *aero.Context) string {
// 	// 	var titles []string

// 	// 	results := make(chan *arn.Anime)
// 	// 	arn.Scan("Anime", results)

// 	// 	for anime := range results {
// 	// 		titles = append(titles, anime.Title.Romaji)
// 	// 	}
// 	// 	sort.Strings(titles)

// 	// 	return ctx.Error(toString(len(titles)) + "\n\n" + strings.Join(titles, "\n"))
// 	// })

// 	app.Get("/api/anime/:id", func(ctx *aero.Context) string {
// 		id := ctx.Get("id")
// 		anime, err := arn.GetAnime(id)

// 		if err != nil {
// 			return ctx.Error(404, "Anime not found", err)
// 		}

// 		return ctx.JSON(anime)
// 	})

// 	app.Get("/api/users/:nick", func(ctx *aero.Context) string {
// 		nick := ctx.Get("nick")
// 		user, err := arn.GetUserByNick(nick)

// 		if err != nil {
// 			return ctx.Error(404, "User not found", err)
// 		}

// 		return ctx.JSON(user)
// 	})

// 	app.Get("/api/threads/:id", func(ctx *aero.Context) string {
// 		id := ctx.Get("id")
// 		thread, err := arn.GetThread(id)

// 		if err != nil {
// 			return ctx.Error(404, "Thread not found", err)
// 		}

// 		return ctx.JSON(thread)
// 	})

// 	animeListAPI := &AnimeListAPI{}
// 	app.Get("/api/anime/:id/add", APIAdd(animeListAPI))
// }

// type AnimeListAPI struct {
// }

// func (api *AnimeListAPI) Key(user *arn.User) interface{} {
// 	return user.ID
// }

// func (api *AnimeListAPI) Table() string {
// 	return "AnimeList"
// }

// func (api *AnimeListAPI) NewList() EasyAPIList {
// 	return new(arn.AnimeList)
// }

// func (api *AnimeListAPI) ListItemID(ctx *aero.Context) string {
// 	return ctx.Get("id")
// }

// // EasyAPI
// type EasyAPI interface {
// 	Key(*arn.User) interface{}
// 	Table() string
// 	NewList() GenericList
// 	ListItemID(*aero.Context) string
// }

// // GenericList ...
// type GenericList interface {
// 	Add(string) error
// 	Save() error
// }

// // APIAdd ...
// func APIAdd(api EasyAPI) aero.Handle {
// 	return func(ctx *aero.Context) string {
// 		objectID := api.ListItemID(ctx)

// 		// Auth
// 		user := utils.GetUser(ctx)

// 		if user == nil {
// 			return ctx.Error(http.StatusBadRequest, "Not logged in", errors.New("User not logged in"))
// 		}

// 		list := api.NewList()

// 		// Fetch list
// 		arn.GetObject(api.Table(), api.Key(user), list)

// 		// Add
// 		addError := list.Add(objectID)

// 		if addError != nil {
// 			return ctx.Error(http.StatusBadRequest, addError.Error(), addError)
// 		}

// 		// Save
// 		saveError := list.Save()

// 		if saveError != nil {
// 			return ctx.Error(http.StatusInternalServerError, "Could not save anime list in database", saveError)
// 		}

// 		// Respond
// 		return ctx.JSON(list)
// 	}
// }

// // func(ctx *aero.Context) string {
// // 	animeID := ctx.Get("id")
// // 	user := utils.GetUser(ctx)

// // if user == nil {
// // 	return ctx.Error(http.StatusBadRequest, "Not logged in", errors.New("User not logged in"))
// // }

// // 	animeList := user.AnimeList()

// // 	// Add
// // 	addError := animeList.Add(animeID)

// // if addError != nil {
// // 	return ctx.Error(http.StatusBadRequest, "Failed adding anime", addError)
// // }

// // 	// Save
// // 	saveError := animeList.Save()

// // if saveError != nil {
// // 	return ctx.Error(http.StatusInternalServerError, "Could not save anime list in database", saveError)
// // }

// // 	return ctx.JSON(animeList)
// // }
