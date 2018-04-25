package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/utils/routetests"
)

func TestRoutes(t *testing.T) {
	t.Parallel()
	app := configure(aero.New())

	// Iterate through every route
	for _, examples := range routetests.All() {
		// Iterate through every example specified for that route
		for _, example := range examples {
			testRoute(t, app, example)
		}
	}
}

func TestAnimePages(t *testing.T) {
	t.Parallel()
	app := configure(aero.New())

	for anime := range arn.StreamAnime() {
		testRoute(t, app, anime.Link())
	}
}

func TestSoundTrackPages(t *testing.T) {
	t.Parallel()
	app := configure(aero.New())

	for soundtrack := range arn.StreamSoundTracks() {
		testRoute(t, app, soundtrack.Link())
	}
}

func TestAMVPages(t *testing.T) {
	t.Parallel()
	app := configure(aero.New())

	for amv := range arn.StreamAMVs() {
		testRoute(t, app, amv.Link())
	}
}

func TestCompanyPages(t *testing.T) {
	t.Parallel()
	app := configure(aero.New())

	for company := range arn.StreamCompanies() {
		testRoute(t, app, company.Link())
	}
}

func TestThreadPages(t *testing.T) {
	t.Parallel()
	app := configure(aero.New())

	for thread := range arn.StreamThreads() {
		testRoute(t, app, thread.Link())
	}
}

func TestPostPages(t *testing.T) {
	t.Parallel()
	app := configure(aero.New())

	for post := range arn.StreamPosts() {
		testRoute(t, app, post.Link())
	}
}

func TestQuotePages(t *testing.T) {
	t.Parallel()
	app := configure(aero.New())

	for quote := range arn.StreamQuotes() {
		testRoute(t, app, quote.Link())
	}
}

func TestUserPages(t *testing.T) {
	t.Parallel()
	app := configure(aero.New())

	for user := range arn.StreamUsers() {
		testRoute(t, app, user.Link())
	}
}

func testRoute(t *testing.T, app *aero.Application, route string) {
	// Create a new HTTP request
	request, err := http.NewRequest("GET", route, nil)

	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	responseRecorder := httptest.NewRecorder()
	app.Handler().ServeHTTP(responseRecorder, request)
	status := responseRecorder.Code

	switch status {
	case 200, 302:
		// 200 and 302 are allowed
	default:
		panic(fmt.Errorf("%s | Wrong status code | %v instead of %v", route, status, http.StatusOK))
	}
}
