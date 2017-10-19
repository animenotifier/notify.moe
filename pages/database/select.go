package database

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/aerogo/mirror"
	"github.com/animenotifier/arn"
)

// QueryResponse ..
type QueryResponse struct {
	Results []interface{} `json:"results"`
}

// Select ...
func Select(ctx *aero.Context) string {
	dataTypeName := ctx.Get("data-type")
	field := ctx.Get("field")
	searchValue := ctx.Get("field-value")

	// Empty values
	if dataTypeName == "+" {
		dataTypeName = ""
	}

	if field == "+" {
		field = ""
	}

	if searchValue == "+" {
		searchValue = ""
	}

	// Check empty parameters
	if dataTypeName == "" || field == "" {
		return ctx.Error(http.StatusBadRequest, "Not enough parameters", nil)
	}

	// Check data type parameter
	_, found := arn.DB.Types()[dataTypeName]

	if !found {
		return ctx.Error(http.StatusBadRequest, "Invalid type", nil)
	}

	response := &QueryResponse{
		Results: []interface{}{},
	}

	stream, err := arn.DB.All(dataTypeName)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching data from the database", err)
	}

	process := func(obj interface{}) {
		_, _, value, _ := mirror.GetField(obj, field)

		if value.String() == searchValue {
			response.Results = append(response.Results, obj)
		}
	}

	switch dataTypeName {
	case "Analytics":
		for obj := range stream.(chan *arn.Analytics) {
			process(obj)
		}
	case "Anime":
		for obj := range stream.(chan *arn.Anime) {
			process(obj)
		}
	case "AnimeList":
		for obj := range stream.(chan *arn.AnimeList) {
			process(obj)
		}
	case "Character":
		for obj := range stream.(chan *arn.Character) {
			process(obj)
		}
	case "Group":
		for obj := range stream.(chan *arn.Group) {
			process(obj)
		}
	case "Post":
		for obj := range stream.(chan *arn.Post) {
			process(obj)
		}
	case "Settings":
		for obj := range stream.(chan *arn.Settings) {
			process(obj)
		}
	case "SoundTrack":
		for obj := range stream.(chan *arn.SoundTrack) {
			process(obj)
		}
	case "Thread":
		for obj := range stream.(chan *arn.Thread) {
			process(obj)
		}
	case "User":
		for obj := range stream.(chan *arn.User) {
			process(obj)
		}
	}

	for _, obj := range response.Results {
		mirror.GetField(obj, field)
	}

	return ctx.JSON(response)
}
