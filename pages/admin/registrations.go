package admin

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// UserRegistrations ...
func UserRegistrations(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	if user.Role != "admin" {
		return ctx.Error(http.StatusUnauthorized, "Not authorized")
	}

	total := 0
	yearInfo := map[int]utils.YearRegistrations{}
	years := []int{}

	for user := range arn.StreamUsers() {
		if user.Registered == "" {
			color.Red("%s %s", user.ID, user.Nick)
			user.Registered = user.LastLogin
			user.Save()
		}

		registered := user.RegisteredTime()
		year := registered.Year()
		yearRegistrations := yearInfo[year]
		yearRegistrations.Total++

		if yearRegistrations.Months == nil {
			yearRegistrations.Months = map[int]int{}
		}

		yearRegistrations.Months[int(registered.Month())]++
		yearInfo[year] = yearRegistrations

		total++
	}

	for year := range yearInfo {
		years = append(years, year)
	}

	sort.Ints(years)

	return ctx.HTML(components.UserRegistrations(total, years, yearInfo))
}
