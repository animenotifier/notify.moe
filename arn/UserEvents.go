package arn

import "github.com/aerogo/aero"

// // EventStreams returns the user's active event streams.
// func (user *User) EventStreams() []*aero.EventStream {
// 	return user.eventStreams
// }

// AddEventStream adds an event stream to the given user.
func (user *User) AddEventStream(stream *aero.EventStream) {
	user.eventStreams.Lock()
	defer user.eventStreams.Unlock()

	user.eventStreams.value = append(user.eventStreams.value, stream)
}

// RemoveEventStream removes an event stream from the given user
// and returns true if it was removed, otherwise false.
func (user *User) RemoveEventStream(stream *aero.EventStream) bool {
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
func (user *User) BroadcastEvent(event *aero.Event) {
	user.eventStreams.Lock()
	defer user.eventStreams.Unlock()

	for _, stream := range user.eventStreams.value {
		// Non-blocking send because we don't know if our listeners are still active.
		select {
		case stream.Events <- event:
		default:
		}
	}
}
