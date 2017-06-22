package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aerogo/aero"
)

func TestRoutes(t *testing.T) {
	expectedStatus := http.StatusOK

	app := configure(aero.New())
	request, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	app.Handler().ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != expectedStatus {
		t.Errorf("Wrong status code: %v instead of %v", status, expectedStatus)
	}
}
