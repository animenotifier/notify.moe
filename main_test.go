package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/utils/routetests"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	t.Parallel()
	app := configure(aero.New())
	app.BindMiddleware()

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
	app.BindMiddleware()

	for anime := range arn.StreamAnime() {
		testRoute(t, app, anime.Link())
	}
}

func TestSoundTrackPages(t *testing.T) {
	t.Parallel()
	app := configure(aero.New())
	app.BindMiddleware()

	for soundtrack := range arn.StreamSoundTracks() {
		testRoute(t, app, soundtrack.Link())
		assert.NotNil(t, soundtrack.Creator())
	}
}

func TestAMVPages(t *testing.T) {
	t.Parallel()
	app := configure(aero.New())
	app.BindMiddleware()

	for amv := range arn.StreamAMVs() {
		testRoute(t, app, amv.Link())
		assert.NotNil(t, amv.Creator())
	}
}

func TestCompanyPages(t *testing.T) {
	t.Parallel()
	app := configure(aero.New())
	app.BindMiddleware()

	for company := range arn.StreamCompanies() {
		testRoute(t, app, company.Link())
	}
}

func TestThreadPages(t *testing.T) {
	t.Parallel()
	app := configure(aero.New())
	app.BindMiddleware()

	for thread := range arn.StreamThreads() {
		testRoute(t, app, thread.Link())
		assert.NotNil(t, thread.Creator())
	}
}

func TestPostPages(t *testing.T) {
	t.Parallel()
	app := configure(aero.New())
	app.BindMiddleware()

	for post := range arn.StreamPosts() {
		testRoute(t, app, post.Link())
		assert.NotNil(t, post.Creator())
	}
}

func TestQuotePages(t *testing.T) {
	t.Parallel()
	app := configure(aero.New())
	app.BindMiddleware()

	for quote := range arn.StreamQuotes() {
		testRoute(t, app, quote.Link())
		assert.NotNil(t, quote.Creator())
	}
}

// func TestUserPages(t *testing.T) {
// 	t.Parallel()
// 	app := configure(aero.New())

// 	for user := range arn.StreamUsers() {
// 		testRoute(t, app, user.Link())
// 	}
// }

func testRoute(t *testing.T, app *aero.Application, route string) {
	request := httptest.NewRequest("GET", strings.ReplaceAll(route, " ", "%20"), nil)
	response := httptest.NewRecorder()
	app.ServeHTTP(response, request)
	status := response.Code

	switch status {
	case http.StatusOK, http.StatusTemporaryRedirect, http.StatusPermanentRedirect:
		// OK
	default:
		panic(fmt.Errorf("%s | Wrong status code | %v instead of %v", route, status, http.StatusOK))
	}
}
