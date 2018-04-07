package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/aerogo/aero"
	"github.com/aerogo/http/client"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/fatih/color"
	"github.com/mssola/user_agent"
)

// UserInfo updates user related information after each request.
func UserInfo() aero.Middleware {
	return func(ctx *aero.Context, next func()) {
		next()

		// Ignore non-HTML requests
		contentType := ctx.Response().Header().Get("Content-Type")

		if !strings.HasPrefix(contentType, "text/html") {
			return
		}

		user := utils.GetUser(ctx)

		// When there's no user logged in, nothing to update
		if user == nil {
			return
		}

		// This works asynchronously so it doesn't block the response
		go updateUserInfo(ctx, user)
	}
}

// Update browser and OS data
func updateUserInfo(ctx *aero.Context, user *arn.User) {
	newIP := ctx.RealIP()
	newUserAgent := ctx.UserAgent()

	if user.UserAgent != newUserAgent {
		user.UserAgent = newUserAgent

		// Parse user agent
		parsed := user_agent.New(user.UserAgent)

		// Browser
		user.Browser.Name, user.Browser.Version = parsed.Browser()

		// OS
		os := parsed.OSInfo()
		user.OS.Name = os.Name
		user.OS.Version = os.Version
	}

	if user.IP != newIP {
		updateUserLocation(user, newIP)
	}

	user.LastSeen = arn.DateTimeUTC()
	user.Save()
}

// Updates the location of the user.
func updateUserLocation(user *arn.User, newIP string) {
	user.IP = newIP

	if arn.APIKeys.IPInfoDB.ID == "" {
		if arn.IsProduction() {
			color.Red("IPInfoDB key not defined")
		}

		return
	}

	locationAPI := "https://api.ipinfodb.com/v3/ip-city/?key=" + arn.APIKeys.IPInfoDB.ID + "&ip=" + user.IP + "&format=json"
	response, err := client.Get(locationAPI).End()

	if err != nil {
		color.Red("Couldn't fetch location data | Error: %s | IP: %s", err.Error(), user.IP)
		return
	}

	if response.StatusCode() != http.StatusOK {
		color.Red("Couldn't fetch location data | Status: %d | IP: %s", response.StatusCode, user.IP)
		return
	}

	newLocation := arn.IPInfoDBLocation{}
	response.Unmarshal(&newLocation)

	if newLocation.CountryName != "-" {
		user.Location.CountryName = newLocation.CountryName
		user.Location.CountryCode = newLocation.CountryCode
		user.Location.Latitude, _ = strconv.ParseFloat(newLocation.Latitude, 64)
		user.Location.Longitude, _ = strconv.ParseFloat(newLocation.Longitude, 64)
		user.Location.CityName = newLocation.CityName
		user.Location.RegionName = newLocation.RegionName
		user.Location.TimeZone = newLocation.TimeZone
		user.Location.ZipCode = newLocation.ZipCode
	}
}
