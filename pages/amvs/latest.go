package amvs

import (
	"sort"

	"github.com/aerogo/aero"
)

// Latest AMVs.
func Latest(ctx aero.Context) error {
	amvs := fetchAll()

	sort.Slice(amvs, func(i, j int) bool {
		return amvs[i].Created > amvs[j].Created
	})

	return render(ctx, amvs)
}
