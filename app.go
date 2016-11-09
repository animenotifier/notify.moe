package main

import (
	"io/ioutil"

	"github.com/aerojs/aero"
)

var app = aero.New()

func main() {
	app.SetStyle(bundledCSS)

	scripts, _ := ioutil.ReadFile("temp/scripts.js")
	js := string(scripts)

	app.Get("/scripts.js", func(ctx *aero.Context) string {
		ctx.SetHeader("Content-Type", "application/javascript")
		return js
	})

	app.Layout = func(ctx *aero.Context, content string) string {
		return Render.Layout(content)
	}

	app.Run()
}
