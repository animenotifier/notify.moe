package admin

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxReports = 80

// ClientErrors shows client-side errors.
func ClientErrors(ctx *aero.Context) string {
	reports := arn.AllClientErrorReports()

	sort.Slice(reports, func(i, j int) bool {
		return reports[i].Created > reports[j].Created
	})

	if len(reports) > maxReports {
		reports = reports[:maxReports]
	}

	return ctx.HTML(components.ClientErrors(reports))
}
