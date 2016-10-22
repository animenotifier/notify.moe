package main

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/Joker/jade"
	"github.com/aerojs/aero"
	"github.com/animenotifier/arn"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasttemplate"
)

func main() {
	arn.Init()

	app := aero.New()
	jade.PrettyOutput = false
	code, _ := jade.ParseFile("pages/anime/anime.pug")
	code = strings.TrimSpace(code)
	code = strings.Replace(code, "{{ ", "{{", -1)
	code = strings.Replace(code, " }}", "}}", -1)
	fmt.Println(code)
	t := fasttemplate.New(code, "{{", "}}")

	app.Get("/", func(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
		ctx.Response.Header.Set("content-type", "text/html;charset=utf-8")
		ctx.SetBodyString("Hello World")
	})

	app.Get("/anime/:id", func(ctx *fasthttp.RequestCtx, params fasthttprouter.Params) {
		id, _ := strconv.Atoi(params.ByName("id"))
		anime, err := arn.GetAnime(id)

		if err != nil {
			ctx.WriteString("Anime not found")
			return
		}

		ctx.Response.Header.Set("content-type", "text/html;charset=utf-8")

		writer := ctx.Response.BodyWriter()
		t.ExecuteFunc(writer, func(w io.Writer, tag string) (int, error) {
			val := reflect.ValueOf(*anime)
			parts := strings.Split(tag, ".")

			for _, part := range parts {
				val = val.FieldByName(part)
			}

			switch val.Kind() {
			case reflect.Int:
				num := strconv.FormatInt(val.Int(), 10)
				return w.Write([]byte(num))
			default:
				return w.Write([]byte(val.String()))
			}
		})
	})

	fmt.Println("Starting server on http://localhost:5000/")

	app.Run()
}
