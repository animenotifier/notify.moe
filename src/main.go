package main

import (
	"fmt"
	"time"

	web "github.com/kataras/iris"
)

func main() {
	InitDatabase()

	web.Get("/", func(ctx *web.Context) {
		ctx.Response.Header.Set("Content-Type", "text/html;charset=utf-8")
		ctx.Write(ctx.Request.URI().String())
	})

	web.Get("/anime/:id", func(ctx *web.Context) {
		start := time.Now()
		id, _ := ctx.ParamInt("id")
		anime := GetAnime(id)

		ctx.Write(anime.Title.Romaji + "\n")
		ctx.Write(anime.Description)
		fmt.Println(time.Since(start))
	})

	web.Listen(":8082")
}
