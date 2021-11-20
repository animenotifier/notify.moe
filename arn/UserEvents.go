package arn

import (
	"github.com/aerogo/aero/event"
)

// AddEventStream adds an event stream to the given user.
func (user *User) AddEventStream(stream *event.Stream) {
	user.eventStreams.Lock()
	defer user.eventStreams.Unlock()

	user.eventStreams.value = append(user.eventStreams.value, stream)
}

// RemoveEventStream removes an event stream from the given user
// and returns true if it was removed, otherwise false.
func (user *User) RemoveEventStream(stream *event.Stream) bool {
	user.eventStreams.Lock()
	defer user.eventStreams.Unlock()

	for index, element := range user.eventStreams.value {
		if element == stream {
			user.eventStreams.value = append(user.eventStreams.value[:index], user.eventStreams.value[index+1:]...)
			return true
		}
	}

	return false
}

// BroadcastEvent sends the given event to all event streams for the given user.
func (user *User) BroadcastEvent(evt *event.Event) {
	user.eventStreams.Lock()
	defer user.eventStreams.Unlock()

	for _, stream := range user.eventStreams.value {
		// Non-blocking send because we don't know if our listeners are still active.
		select {
		case stream.Events <- evt:
		default:
		}
	}
}
