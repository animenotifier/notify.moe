package sse

import (
	"fmt"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/utils"
)

// Events streams server events to the client.
func Events(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	fmt.Println(user.Nick, "receiving live events")

	events := make(chan *aero.Event)
	disconnected := make(chan struct{})

	go func() {
		defer fmt.Println(user.Nick, "disconnected, stop sending events")

		for {
			select {
			case <-disconnected:
				close(events)
				return

				// case <-time.After(10 * time.Second):
				// 	events <- &aero.Event{
				// 		Name: "ping",
				// 	}
			}
		}
	}()

	return ctx.EventStream(events, disconnected)
}
