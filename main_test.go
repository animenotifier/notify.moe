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
		if anime.IsDraft {
			continue
		}

		assert.True(t, anime.Title.Canonical != "")
		fetch(t, app, anime.Link())
	}
}

func TestSoundTracks(t *testing.T) {
	app := server.New()
	app.BindMiddleware()

	for soundtrack := range arn.StreamSoundTracks() {
		assert.NotNil(t, soundtrack.Creator())
		fetch(t, app, soundtrack.Link())
	}
}

func TestAMVs(t *testing.T) {
	app := server.New()
	app.BindMiddleware()

	for amv := range arn.StreamAMVs() {
		assert.NotNil(t, amv.Creator())
		fetch(t, app, amv.Link())
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
		assert.NotNil(t, thread.Creator())
		fetch(t, app, thread.Link())
	}
}

func TestPosts(t *testing.T) {
	app := server.New()
	app.BindMiddleware()

	for post := range arn.StreamPosts() {
		assert.NotNil(t, post.Creator())
		fetch(t, app, post.Link())
	}
}

func TestQuotes(t *testing.T) {
	app := server.New()
	app.BindMiddleware()

	for quote := range arn.StreamQuotes() {
		assert.NotNil(t, quote.Creator())
		fetch(t, app, quote.Link())
	}
}

func TestUsers(t *testing.T) {
	app := server.New()
	app.BindMiddleware()

	for user := range arn.StreamUsers() {
		assert.True(t, user.Nick != "")
		// fetch(t, app, user.Link())
	}
}

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
