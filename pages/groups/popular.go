package groups

import (
	"sort"

	"github.com/aerogo/aero"
)

// Popular shows the most popular groups.
func Popular(ctx aero.Context) error {
	groups := fetchGroups("")

	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i].Members) == len(groups[j].Members) {
			return groups[i].Created > groups[j].Created
		}

		return len(groups[i].Members) > len(groups[j].Members)
	})

	return render(ctx, groups)
}
