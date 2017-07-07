package statistics

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Anime ...
func Anime(ctx *aero.Context) string {
	statistics := arn.StatisticsCategory{}
	arn.DB.GetObject("Cache", "anime statistics", &statistics)
	return ctx.HTML(components.Statistics(statistics.PieCharts...))
}
