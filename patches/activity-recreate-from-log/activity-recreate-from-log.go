package main

import "github.com/animenotifier/notify.moe/arn"

func main() {
	defer arn.Node.Close()

	for entry := range arn.StreamEditLogEntries() {
		if entry.Action != "create" {
			continue
		}

		obj := entry.Object()

		if obj == nil {
			continue
		}

		draft, isDraftable := obj.(arn.Draftable)

		if isDraftable && draft.GetIsDraft() {
			continue
		}

		// We don't create activity entries for anything
		// other than posts and threads.
		if entry.ObjectType != "Post" && entry.ObjectType != "Thread" {
			continue
		}

		activity := arn.NewActivityCreate(
			entry.ObjectType,
			entry.ObjectID,
			entry.UserID,
		)

		activity.Created = entry.Created
		activity.Save()
	}
}
