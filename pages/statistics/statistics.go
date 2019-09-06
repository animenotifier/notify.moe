package statistics

import (
	"fmt"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

type stats map[string]float64

// Get ...
func Get(ctx aero.Context) error {
	pieCharts := getUserStats()
	return ctx.HTML(components.Statistics(pieCharts))
}

func getUserStats() []*arn.PieChart {
	screenSize := stats{}
	pixelRatio := stats{}
	browser := stats{}
	country := stats{}
	gender := stats{}
	os := stats{}
	notifications := stats{}
	titleLanguage := stats{}
	ip := stats{}
	pro := stats{}
	connectionType := stats{}
	roundTripTime := stats{}
	downLink := stats{}
	theme := stats{}

	for info := range arn.StreamAnalytics() {
		user, err := arn.GetUser(info.UserID)
		arn.PanicOnError(err)

		if !user.IsActive() {
			continue
		}

		pixelRatio[fmt.Sprintf("%.0f", info.Screen.PixelRatio)]++

		size := fmt.Sprint(info.Screen.Width) + " x " + fmt.Sprint(info.Screen.Height)
		screenSize[size]++

		if info.Connection.EffectiveType != "" {
			connectionType[info.Connection.EffectiveType]++
		}

		if info.Connection.DownLink != 0 {
			downLink[fmt.Sprintf("%.0f Mb/s", info.Connection.DownLink)]++
		}

		if info.Connection.RoundTripTime != 0 {
			roundTripTime[fmt.Sprintf("%.0f ms", info.Connection.RoundTripTime)]++
		}
	}

	for user := range arn.StreamUsers() {
		if !user.IsActive() {
			continue
		}

		if user.Gender != "" && user.Gender != "other" {
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

		if len(user.PushSubscriptions().Items) > 0 {
			notifications["Enabled"]++
		} else {
			notifications["Disabled"]++
		}

		if arn.IsIPv6(user.IP) {
			ip["IPv6"]++
		} else {
			ip["IPv4"]++
		}

		if user.IsPro() {
			pro["PRO accounts"]++
		} else {
			pro["Free accounts"]++
		}

		settings := user.Settings()
		theme[settings.Theme]++
		titleLanguage[settings.TitleLanguage]++
	}

	return []*arn.PieChart{
		arn.NewPieChart("OS", os),
		arn.NewPieChart("Screen size", screenSize),
		arn.NewPieChart("Browser", browser),
		arn.NewPieChart("Country", country),
		arn.NewPieChart("Title language", titleLanguage),
		arn.NewPieChart("Notifications", notifications),
		arn.NewPieChart("Gender", gender),
		arn.NewPieChart("Pixel ratio", pixelRatio),
		arn.NewPieChart("Download speed", downLink),
		arn.NewPieChart("Ping", roundTripTime),
		arn.NewPieChart("Connection", connectionType),
		arn.NewPieChart("IP version", ip),
		arn.NewPieChart("PRO accounts", pro),
		arn.NewPieChart("Theme", theme),
	}
}
