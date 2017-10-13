package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aerogo/aero"
	"github.com/fatih/color"
)

func TestRoutes(t *testing.T) {
	app := configure(aero.New())

	for _, examples := range routeTests {
		for _, example := range examples {
			request, err := http.NewRequest("GET", example, nil)

			if err != nil {
				t.Fatal(err)
			}

			responseRecorder := httptest.NewRecorder()
			app.Handler().ServeHTTP(responseRecorder, request)

			if status := responseRecorder.Code; status != http.StatusOK {
				color.Red("%s | Wrong status code | %v instead of %v", example, status, http.StatusOK)
			}
		}
	}
}
