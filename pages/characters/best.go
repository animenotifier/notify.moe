package characters

import (
	"sort"

	"github.com/aerogo/aero"
)

// Best characters.
func Best(ctx *aero.Context) string {
	characters := fetchAll()

	sort.Slice(characters, func(i, j int) bool {
		if len(characters[i].Likes) == len(characters[j].Likes) {
			return characters[i].Name.Canonical < characters[j].Name.Canonical
		}

		return len(characters[i].Likes) > len(characters[j].Likes)
	})

	return render(ctx, characters)
}
