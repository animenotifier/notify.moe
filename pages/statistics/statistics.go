package statistics

import (
	"github.com/aerogo/aero"
)

// Get ...
func Get(ctx *aero.Context) string {
	// statistics := arn.StatisticsCategory{}
	// arn.DB.GetObject("Cache", "user statistics", &statistics)
	// return ctx.HTML(components.Statistics(statistics.PieCharts...))
	return ctx.HTML("Not implemented")
}
