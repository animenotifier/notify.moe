package main

import (
	"fmt"

	"github.com/aerogo/nano"
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Deleting old sessions")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	// threshold := time.Now().Add(-6 * 30 * 24 * time.Hour).Format(time.RFC3339)
	count := 0
	total := 0

	for session := range streamSessions() {
		data := *session
		// created := data["created"].(string)

		if data["userId"] == nil { // created < threshold
			arn.DB.Delete("Session", data["sid"].(string))
			count++
		}

		total++
	}

	fmt.Printf("Deleted %d / %d sessions.\n", count, total)
}

func streamSessions() chan *arn.Session {
	channel := make(chan *arn.Session, nano.ChannelBufferSize)

	go func() {
		for obj := range arn.DB.All("Session") {
			channel <- obj.(*arn.Session)
		}

		close(channel)
	}()

	return channel
}
