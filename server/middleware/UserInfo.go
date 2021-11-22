package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/aerogo/aero"
	"github.com/aerogo/http/client"
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/mssola/user_agent"
)

// UserInfo updates user related information after each request.
func UserInfo(next aero.Handler) aero.Handler {
	return func(ctx aero.Context) error {
		err := next(ctx)

		// Ignore non-HTML requests
		contentType := ctx.Response().Header("Content-Type")

		if !strings.HasPrefix(contentType, "text/html") {
			return nil
		}

		user := arn.GetUserFromContext(ctx)

		// When there's no user logged in, nothing to update
		if user == nil {
			return nil
		}

		// Bind local variables and start a coroutine
		ip := ctx.IP()
		userAgent := ctx.Request().Header("User-Agent")
		go updateUserInfo(ip, userAgent, user)

		return err
	}
}

// Update browser and OS data
func updateUserInfo(ip string, userAgent string, user *arn.User) {
	if user.UserAgent != userAgent {
		user.UserAgent = userAgent

		// Parse user agent
		parsed := user_agent.New(user.UserAgent)

		// Browser
		user.Browser.Name, user.Browser.Version = parsed.Browser()

		// OS
		os := parsed.OSInfo()
		user.OS.Name = os.Name
		user.OS.Version = os.Version
	}

	user.LastSeen = arn.DateTimeUTC()
	user.Save()

	if user.IP == ip {
		return
	}

	updateUserLocation(user, ip)
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
	err = response.Unmarshal(&newLocation)

	if err != nil {
		color.Red("Couldn't deserialize location data | Status: %d | IP: %s", response.StatusCode, user.IP)
		return
	}

	if newLocation.CountryName == "" || newLocation.CountryName == "-" {
		return
	}

	user.Location.CountryName = newLocation.CountryName
	user.Location.CountryCode = newLocation.CountryCode
	user.Location.Latitude, _ = strconv.ParseFloat(newLocation.Latitude, 64)
	user.Location.Longitude, _ = strconv.ParseFloat(newLocation.Longitude, 64)
	user.Location.CityName = newLocation.CityName
	user.Location.RegionName = newLocation.RegionName
	user.Location.TimeZone = newLocation.TimeZone
	user.Location.ZipCode = newLocation.ZipCode

	// Make South Korea easier to read
	if user.Location.CountryName == "Korea, Republic of" || user.Location.CountryName == "Korea (Republic of)" {
		user.Location.CountryName = "South Korea"
	}

	user.Save()
}
