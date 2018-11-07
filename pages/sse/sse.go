package sse

import (
	"fmt"
	"net/http"

	"github.com/animenotifier/notify.moe/components/css"
	"github.com/animenotifier/notify.moe/components/js"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/utils"
)

var (
	scriptsETag = aero.ETagString(js.Bundle())
	stylesETag  = aero.ETagString(css.Bundle())
	streams     = map[string][]*aero.EventStream{}
)

// Events streams server events to the client.
func Events(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	fmt.Println(user.Nick, "receiving live events")
	stream := aero.NewEventStream()
	user.AddEventStream(stream)

	go func() {
		defer fmt.Println(user.Nick, "disconnected, stop sending events")

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

		for {
			select {
			case <-stream.Closed:
				user.RemoveEventStream(stream)
				return

				// case <-time.After(10 * time.Second):
				// 	stream.Events <- &aero.Event{
				// 		Name: "ping",
				// 	}
			}
		}
	}()

	return ctx.EventStream(stream)
}
