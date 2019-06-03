package groups

import "github.com/animenotifier/notify.moe/arn"

func fetchGroups(memberID string) []*arn.Group {
	return arn.FilterGroups(func(group *arn.Group) bool {
		if group.IsDraft {
			return false
		}

		if memberID != "" && !group.HasMember(memberID) {
			return false
		}

		return true
	})
}
