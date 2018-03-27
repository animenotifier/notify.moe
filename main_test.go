package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/utils/routetests"
)

// TestRouteStatusCodes tests the status code of every route registered in routeTests.
func TestRouteStatusCodes(t *testing.T) {
	app := configure(aero.New())

	// Iterate through every route
	for _, examples := range routetests.All() {
		// Iterate through every example specified for that route
		for _, example := range examples {
			// Create a new HTTP request
			request, err := http.NewRequest("GET", example, nil)

			if err != nil {
				t.Fatal(err)
			}

			// Record the response without actually starting the server
			responseRecorder := httptest.NewRecorder()
			app.Handler().ServeHTTP(responseRecorder, request)

			if status := responseRecorder.Code; status != http.StatusOK {
				panic(fmt.Errorf("%s | Wrong status code | %v instead of %v", example, status, http.StatusOK))
			}
		}
	}
}
