package statistics

import (
	"fmt"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get ...
func Get(ctx *aero.Context) string {
	analytics, err := arn.AllAnalytics()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Couldn't retrieve analytics", err)
	}

	screenSize := map[string]float64{}
	platform := map[string]float64{}
	pixelRatio := map[string]float64{}

	for _, info := range analytics {
		platform[info.System.Platform]++
		pixelRatio[fmt.Sprintf("%.1f", info.Screen.PixelRatio)]++

		size := arn.ToString(info.Screen.Width) + " x " + arn.ToString(info.Screen.Height)
		screenSize[size]++
	}

	return ctx.HTML(components.Statistics(
		utils.NewPieChart("Screen sizes", screenSize),
		utils.NewPieChart("Platforms", platform),
		utils.NewPieChart("Pixel ratios", pixelRatio),
	))
}
