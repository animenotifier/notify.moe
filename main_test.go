package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// func BenchmarkRender(b *testing.B) {
// 	b.ReportAllocs()
// 	layout := aero.NewTemplate("layout/layout.pug")
// 	template := aero.NewTemplate("pages/anime/anime.pug")
// 	cssBytes, _ := ioutil.ReadFile("layout.css")
// 	css := string(cssBytes)
// 	cssSearch := "</head><body"
// 	cssReplace := "<style>" + css + "</style></head><body"

// 	for i := 0; i < b.N; i++ {
// 		anime, _ := arn.GetAnime(1000001)

// 		content := template.Render(map[string]interface{}{
// 			"anime": anime,
// 		})

// 		final := layout.Render(map[string]interface{}{
// 			"content": content,
// 		})

// 		final = strings.Replace(final, cssSearch, cssReplace, 1)
// 	}
// }

func Benchmark1(b *testing.B) {
	code := []byte(strings.Repeat("a", 10000))

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		buf.Write(code)
		buf.Write(code)
		buf.Write(code)
		c := buf.String()
		fmt.Println(len(c))
	}
}

func Benchmark2(b *testing.B) {
	code := strings.Repeat("a", 10000)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c := code
		c += code
		c += code
		fmt.Println(len(c))
	}
}
