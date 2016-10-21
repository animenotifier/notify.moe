package main

import (
	"fmt"
	"time"

	"github.com/kataras/go-template/pug"
	"github.com/kataras/iris"
)

func main() {
	InitDatabase()

	iris.Config.Gzip = true
	iris.Config.IsDevelopment = true

	cfg := pug.DefaultConfig()
	cfg.Layout = "layout.pug"

	iris.UseTemplate(pug.New(cfg)).Directory("pages", ".pug")

	iris.Static("/styles", "./styles", 1)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Response.Header.Set("Content-Type", "text/html;charset=utf-8")
		ctx.Write(ctx.Request.URI().String())
	})

	iris.Get("/anime/:id", func(ctx *iris.Context) {
		start := time.Now()
		id, _ := ctx.ParamInt("id")
		anime := GetAnime(id)

		ctx.MustRender("anime.pug", anime)
		fmt.Println(time.Since(start))
	})

	iris.Listen(":8082")
}
