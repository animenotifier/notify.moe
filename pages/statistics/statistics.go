package statistics

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

type stats = map[string]float64

// Get ...
func Get(ctx *aero.Context) string {
	analytics, err := arn.AllAnalytics()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Couldn't retrieve analytics", err)
	}

	screenSize := stats{}
	// platform := stats{}
	pixelRatio := stats{}
	browser := stats{}
	country := stats{}
	gender := stats{}
	os := stats{}

	for _, info := range analytics {
		// platform[info.System.Platform]++
		pixelRatio[fmt.Sprintf("%.1f", info.Screen.PixelRatio)]++

		size := arn.ToString(info.Screen.Width) + " x " + arn.ToString(info.Screen.Height)
		screenSize[size]++
	}

	for user := range arn.MustStreamUsers() {
		if user.Gender != "" {
			gender[user.Gender]++
		}

		if user.Browser.Name != "" {
			browser[user.Browser.Name]++
		}

		if user.Location.CountryName != "" {
			country[user.Location.CountryName]++
		}

		if user.OS.Name != "" {
			if strings.HasPrefix(user.OS.Name, "CrOS") {
				user.OS.Name = "Chrome OS"
			}

			os[user.OS.Name]++
		}
	}

	return ctx.HTML(components.Statistics(
		utils.NewPieChart("OS", os),
		// utils.NewPieChart("Platform", platform),
		utils.NewPieChart("Screen size", screenSize),
		utils.NewPieChart("Pixel ratio", pixelRatio),
		utils.NewPieChart("Browser", browser),
		utils.NewPieChart("Country", country),
		utils.NewPieChart("Gender", gender),
	))
}
