package arn

import (
	"github.com/aerogo/nano"
)

// UserFollows is a list including IDs to users you follow.
type UserFollows struct {
	UserID UserID   `json:"userId" primary:"true"`
	Items  []string `json:"items"`
}

// StreamUserFollows returns a stream of all user follows.
func StreamUserFollows() <-chan *UserFollows {
	channel := make(chan *UserFollows, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("UserFollows") {
			channel <- obj.(*UserFollows)
		}

		close(channel)
	}()

	return channel
}
