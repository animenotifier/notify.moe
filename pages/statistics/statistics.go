package statistics

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	statistics := arn.StatisticsCategory{}
	arn.DB.GetObject("Cache", "user statistics", &statistics)
	return ctx.HTML(components.Statistics(statistics.PieCharts...))
}
