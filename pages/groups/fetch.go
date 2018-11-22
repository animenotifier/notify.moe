package groups

import "github.com/animenotifier/arn"

func fetchGroups() []*arn.Group {
	return arn.FilterGroups(func(group *arn.Group) bool {
		return !group.IsDraft
	})
}
