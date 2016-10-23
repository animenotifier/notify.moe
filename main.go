package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/Joker/jade"
	"github.com/OneOfOne/xxhash"
	"github.com/aerojs/aero"
	"github.com/blitzprog/arn"
	"github.com/buaazp/fasthttprouter"
	cache "github.com/patrickmn/go-cache"
	"github.com/robertkrimen/otto"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasttemplate"
)

const (
	gzipThreshold         = 1450
	contentTypeHeader     = "content-type"
	contentType           = "text/html;charset=utf-8"
	contentEncodingHeader = "content-encoding"
	contentEncoding       = "gzip"
	hello                 = "Hello World"
)

var jsonToResponse *cache.Cache

func worker(script *otto.Script, jobs <-chan map[string]interface{}, results chan<- string) {
	vm := otto.New()

	for properties := range jobs {
		h := xxhash.NewS64(0)

		for _, value := range properties {
			jsonBytes, err := json.Marshal(value)

			if err == nil {
				h.Write(jsonBytes)
				// fmt.Println(string(buffer))
			}
		}

		hash := strconv.FormatUint(h.Sum64(), 10)
		cachedResponse, found := jsonToResponse.Get(hash)

		if found {
			results <- cachedResponse.(string)
		} else {
			for key, value := range properties {
				vm.Set(key, value)
			}

			result, _ := vm.Run(script)
			code, _ := result.ToString()
			results <- code

			jsonToResponse.Set(hash, code, cache.DefaultExpiration)
		}
	}
}

func main() {
	jsonToResponse = cache.New(5*time.Minute, 1*time.Minute)
	app := aero.New()
	jade.PrettyOutput = false
	code, _ := jade.ParseFile("pages/anime/anime.pug")
	code = strings.TrimSpace(code)
	code = strings.Replace(code, "{{ ", "{{", -1)
	code = strings.Replace(code, " }}", "}}", -1)
	code = strings.Replace(code, "\n", " ", -1)
	t, _ := fasttemplate.NewTemplate(code, "{{", "}}")

	jsCode := "html = '" + t.ExecuteFuncString(func(w io.Writer, tag string) (int, error) {
		if tag == "end" {
			return w.Write([]byte("'; }\nhtml += '"))
		}

		if strings.HasPrefix(tag, "if ") {
			return w.Write([]byte("';\nif(" + tag[3:] + ") { html += '"))
		}

		if tag == "else" {
			return w.Write([]byte("';\nelse { html += '"))
		}

		return w.Write([]byte("';\nhtml += (" + tag + ");\nhtml += '"))
	}) + "';"

	jsCode = strings.Replace(jsCode, "html += '';", "", -1)

	// fmt.Println(code)
	// fmt.Println(jsCode)

	jsCompiler := otto.New()
	script, err := jsCompiler.Compile("pages/anime/anime.pug", jsCode)

	if err != nil {
		panic(err)
	}

	example, _ := ioutil.ReadFile("security/frontpage.html")

	jobs := make(chan map[string]interface{}, 4096)
	results := make(chan string, 4096)

	for w := 1; w <= runtime.NumCPU(); w++ {
		go worker(script, jobs, results)
	}

	app.Get("/", func(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
		aero.RespondBytes(ctx, example)
	})

	app.Get("/hello", func(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
		aero.Respond(ctx, "Hello World")
	})

	app.Get("/anime/:id", func(ctx *fasthttp.RequestCtx, params fasthttprouter.Params) {
		id, _ := strconv.Atoi(params.ByName("id"))
		anime, err := arn.GetAnime(id)

		if err != nil {
			aero.Respond(ctx, "Anime not found")
			return
		}

		templateParams := make(map[string]interface{})
		templateParams["anime"] = anime
		jobs <- templateParams

		aero.Respond(ctx, <-results)
	})

	fmt.Println("Starting server on http://localhost:4000/")

	app.Run()
}
