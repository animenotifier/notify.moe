package main

import (
	"fmt"
	"strings"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

type stats map[string]float64

func main() {
	color.Yellow("Generating statistics")

	userStats := getUserStats()
	animeStats := getAnimeStats()

	arn.PanicOnError(arn.DB.Set("Cache", "user statistics", &arn.StatisticsCategory{
		Name:      "Users",
		PieCharts: userStats,
	}))
	arn.PanicOnError(arn.DB.Set("Cache", "anime statistics", &arn.StatisticsCategory{
		Name:      "Anime",
		PieCharts: animeStats,
	}))

	color.Green("Finished.")
}

func getUserStats() []*arn.PieChart {
	println("Generating user statistics")

	analytics, err := arn.AllAnalytics()
	arn.PanicOnError(err)

	screenSize := stats{}
	pixelRatio := stats{}
	browser := stats{}
	country := stats{}
	gender := stats{}
	os := stats{}

	for _, info := range analytics {
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

	println("Finished user statistics")

	return []*arn.PieChart{
		arn.NewPieChart("OS", os),
		arn.NewPieChart("Screen size", screenSize),
		arn.NewPieChart("Browser", browser),
		arn.NewPieChart("Country", country),
		arn.NewPieChart("Gender", gender),
		arn.NewPieChart("Pixel ratio", pixelRatio),
	}
}

func getAnimeStats() []*arn.PieChart {
	println("Generating anime statistics")

	allAnime, err := arn.AllAnime()
	arn.PanicOnError(err)

	shoboi := stats{}
	anilist := stats{}
	status := stats{}
	types := stats{}
	shoboiEdits := stats{}
	anilistEdits := stats{}
	rating := stats{}

	for _, anime := range allAnime {
		for _, external := range anime.Mappings {
			if external.Service == "shoboi/anime" {
				if external.CreatedBy == "" {
					shoboiEdits["Bot"]++
				} else {
					user, err := arn.GetUser(external.CreatedBy)
					arn.PanicOnError(err)
					shoboiEdits[user.Nick]++
				}
			}

			if external.Service == "anilist/anime" {
				if external.CreatedBy == "" {
					anilistEdits["Bot"]++
				} else {
					user, err := arn.GetUser(external.CreatedBy)
					arn.PanicOnError(err)
					anilistEdits[user.Nick]++
				}
			}
		}

		if anime.GetMapping("shoboi/anime") != "" {
			shoboi["Connected with Shoboi"]++
		} else {
			shoboi["Not connected with Shoboi"]++
		}

		if anime.GetMapping("anilist/anime") != "" {
			anilist["Connected with Anilist"]++
		} else {
			anilist["Not connected with Anilist"]++
		}

		rating[arn.ToString(int(anime.Rating.Overall+0.5))]++

		status[anime.Status]++
		types[anime.Type]++
	}

	println("Finished anime statistics")

	return []*arn.PieChart{
		arn.NewPieChart("Type", types),
		arn.NewPieChart("Status", status),
		arn.NewPieChart("Rating", rating),
		arn.NewPieChart("Anilist", anilist),
		arn.NewPieChart("Shoboi", shoboi),
		arn.NewPieChart("Anilist Editors", anilistEdits),
		arn.NewPieChart("Shoboi Editors", shoboiEdits),
	}
}
