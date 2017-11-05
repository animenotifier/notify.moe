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

	stream := arn.DB.All(dataTypeName)

	process := func(obj interface{}) {
		_, _, value, _ := mirror.GetField(obj, field)

		if value.String() == searchValue {
			response.Results = append(response.Results, obj)
		}
	}

	for obj := range stream {
		process(obj)
	}

	for _, obj := range response.Results {
		mirror.GetField(obj, field)
	}

	return ctx.JSON(response)
}
