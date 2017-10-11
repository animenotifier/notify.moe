package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
	"github.com/animenotifier/arn"
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

func TestInterfaceImplementations(t *testing.T) {
	// API interfaces
	var creatable = reflect.TypeOf((*api.Creatable)(nil)).Elem()
	var editable = reflect.TypeOf((*api.Editable)(nil)).Elem()
	var actionable = reflect.TypeOf((*api.Actionable)(nil)).Elem()
	var collection = reflect.TypeOf((*api.Collection)(nil)).Elem()

	// Required interface implementations
	var interfaceImplementations = map[string][]reflect.Type{
		"User": []reflect.Type{
			editable,
		},
		"Thread": []reflect.Type{
			creatable,
			editable,
			actionable,
		},
		"Post": []reflect.Type{
			creatable,
			editable,
			actionable,
		},
		"SoundTrack": []reflect.Type{
			creatable,
		},
		"Analytics": []reflect.Type{
			creatable,
		},
		"AnimeList": []reflect.Type{
			collection,
		},
		"PushSubscriptions": []reflect.Type{
			collection,
		},
		"UserFollows": []reflect.Type{
			collection,
		},
	}

	for typeName, interfaces := range interfaceImplementations {
		for _, requiredInterface := range interfaces {
			if !reflect.PtrTo(arn.DB.Type(typeName)).Implements(requiredInterface) {
				panic(errors.New(typeName + " does not implement interface " + requiredInterface.Name()))
			}
		}
	}
}
