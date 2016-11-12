package main

import (
	"io/ioutil"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
)

var app = aero.New()

func main() {
	// app.SetStyle(bundledCSS)
	app.SetStyle("")

	scripts, _ := ioutil.ReadFile("temp/scripts.js")
	js := string(scripts)

	app.Get("/scripts.js", func(ctx *aero.Context) string {
		ctx.SetHeader("Content-Type", "application/javascript")
		return js
	})

	app.Get("/hello", func(ctx *aero.Context) string {
		return "Hello World"
	})

	app.Layout = func(ctx *aero.Context, content string) string {
		return components.Layout(content)
	}

	app.Run()
}
