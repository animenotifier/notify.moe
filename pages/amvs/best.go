package amvs

import (
	"sort"

	"github.com/aerogo/aero"
)

// Best AMVs.
func Best(ctx aero.Context) error {
	amvs := fetchAll()

	sort.Slice(amvs, func(i, j int) bool {
		if len(amvs[i].Likes) == len(amvs[j].Likes) {
			return amvs[i].Title.String() < amvs[j].Title.String()
		}

		return len(amvs[i].Likes) > len(amvs[j].Likes)
	})

	return render(ctx, amvs)
}
