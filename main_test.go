package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/akyoto/assert"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/server"
	"github.com/animenotifier/notify.moe/utils/routetests"
)

func TestRoutes(t *testing.T) {
	app := server.New()
	app.BindMiddleware()

	// Iterate through every route
	for _, examples := range routetests.All() {
		// Iterate through every example specified for that route
		for _, example := range examples {
			fetch(t, app, example)
		}
	}
}

func TestAnime(t *testing.T) {
	app := server.New()
	app.BindMiddleware()

	for anime := range arn.StreamAnime() {
		fetch(t, app, anime.Link())
	}
}

func TestSoundTracks(t *testing.T) {
	app := server.New()
	app.BindMiddleware()

	for soundtrack := range arn.StreamSoundTracks() {
		fetch(t, app, soundtrack.Link())
		assert.NotNil(t, soundtrack.Creator())
	}
}

func TestAMVs(t *testing.T) {
	app := server.New()
	app.BindMiddleware()

	for amv := range arn.StreamAMVs() {
		fetch(t, app, amv.Link())
		assert.NotNil(t, amv.Creator())
	}
}

func TestCompanies(t *testing.T) {
	app := server.New()
	app.BindMiddleware()

	for company := range arn.StreamCompanies() {
		fetch(t, app, company.Link())
	}
}

func TestThreads(t *testing.T) {
	app := server.New()
	app.BindMiddleware()

	for thread := range arn.StreamThreads() {
		fetch(t, app, thread.Link())
		assert.NotNil(t, thread.Creator())
	}
}

func TestPosts(t *testing.T) {
	app := server.New()
	app.BindMiddleware()

	for post := range arn.StreamPosts() {
		fetch(t, app, post.Link())
		assert.NotNil(t, post.Creator())
	}
}

func TestQuotes(t *testing.T) {
	app := server.New()
	app.BindMiddleware()

	for quote := range arn.StreamQuotes() {
		fetch(t, app, quote.Link())
		assert.NotNil(t, quote.Creator())
	}
}

// func TestUsers(t *testing.T) {
// 	app := server.New()
// 	app.BindMiddleware()

// 	for user := range arn.StreamUsers() {
// 		fetch(t, app, user.Link())
// 	}
// }

func fetch(t *testing.T, app http.Handler, route string) {
	request := httptest.NewRequest("GET", strings.ReplaceAll(route, " ", "%20"), nil)
	response := httptest.NewRecorder()
	app.ServeHTTP(response, request)
	status := response.Code

	switch status {
	case http.StatusOK, http.StatusTemporaryRedirect, http.StatusPermanentRedirect:
		// OK
	default:
		t.Fatalf("%s | Wrong status code | %v instead of %v", route, status, http.StatusOK)
	}
}
