package main

import (
	"github.com/aerogo/api"
	"github.com/animenotifier/arn"
)

func init() {
	api := api.New("/api/", arn.DB)
	api.Install(app)
}
