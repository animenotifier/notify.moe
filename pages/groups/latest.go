package groups

import (
	"sort"

	"github.com/aerogo/aero"
)

// Latest shows the latest groups.
func Latest(ctx aero.Context) error {
	groups := fetchGroups("")

	sort.Slice(groups, func(i, j int) bool {
		return groups[i].Created > groups[j].Created
	})

	return render(ctx, groups)
}
