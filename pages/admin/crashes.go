package admin

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxCrashes = 80

// Crashes shows client-side errors.
func Crashes(ctx aero.Context) error {
	crashes := arn.AllCrashes()

	sort.Slice(crashes, func(i, j int) bool {
		return crashes[i].Created > crashes[j].Created
	})

	if len(crashes) > maxCrashes {
		crashes = crashes[:maxCrashes]
	}

	return ctx.HTML(components.Crashes(crashes))
}
