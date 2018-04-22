package characters

import (
	"sort"

	"github.com/aerogo/aero"
)

// Latest characters.
func Latest(ctx *aero.Context) string {
	characters := fetchAll()

	sort.Slice(characters, func(i, j int) bool {
		return characters[i].Created > characters[j].Created
	})

	return render(ctx, characters)
}
