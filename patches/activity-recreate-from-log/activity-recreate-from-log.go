package main

import "github.com/animenotifier/arn"

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

		draft, isDraftable := obj.(arn.HasDraft)

		if isDraftable && draft.IsDraft {
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
