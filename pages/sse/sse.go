package sse

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components/css"
	"github.com/animenotifier/notify.moe/components/js"
)

var (
	scriptsETag = aero.ETagString(js.Bundle())
	stylesETag  = aero.ETagString(css.Bundle())
)

// Events streams server events to the client.
func Events(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	stream := aero.NewEventStream()
	user.AddEventStream(stream)

	go func() {
		defer user.RemoveEventStream(stream)

		// Send the ETag for the scripts
		stream.Events <- &aero.Event{
			Name: "etag",
			Data: struct {
				URL  string `json:"url"`
				ETag string `json:"etag"`
			}{
				URL:  "/scripts",
				ETag: scriptsETag,
			},
		}

		// Send the ETag for the styles
		stream.Events <- &aero.Event{
			Name: "etag",
			Data: struct {
				URL  string `json:"url"`
				ETag string `json:"etag"`
			}{
				URL:  "/styles",
				ETag: stylesETag,
			},
		}

		// Wait until the user closes the tab or disconnects
		<-stream.Closed
	}()

	return ctx.EventStream(stream)
}
