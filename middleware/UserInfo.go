package middleware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/fatih/color"
	"github.com/mssola/user_agent"
	"github.com/parnurzeal/gorequest"
)

var apiKeys arn.APIKeys

func init() {
	data, _ := ioutil.ReadFile("security/api-keys.json")
	err := json.Unmarshal(data, &apiKeys)

	if err != nil {
		panic(err)
	}
}

// UserInfo updates user related information after each request.
func UserInfo() aero.Middleware {
	return func(ctx *aero.Context, next func()) {
		next()

		// This works asynchronously so it doesn't block the response
		go updateUserInfo(ctx)
	}
}

// updateUserInfo is started asynchronously so it doesn't block the request
func updateUserInfo(ctx *aero.Context) {
	user := utils.GetUser(ctx)

	// When there's no user logged in, nothing to update
	if user == nil {
		return
	}

	// Ignore non-HTML requests
	if strings.Index(ctx.GetRequestHeader("Accept"), "text/html") == -1 {
		return
	}

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
		user.IP = newIP
		locationAPI := "https://api.ipinfodb.com/v3/ip-city/?key=" + apiKeys.IPInfoDB.ID + "&ip=" + user.IP + "&format=json"

		response, data, err := gorequest.New().Get(locationAPI).EndBytes()

		if len(err) > 0 && err[0] != nil {
			color.Red("Couldn't fetch location data | Error: %s | IP: %s", err[0].Error(), user.IP)
			return
		}

		if response.StatusCode != http.StatusOK {
			color.Red("Couldn't fetch location data | Status: %d | IP: %s", response.StatusCode, user.IP)
			return
		}

		newLocation := arn.IPInfoDBLocation{}
		json.Unmarshal(data, &newLocation)

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

	user.LastSeen = arn.DateTimeUTC()
	user.Save()
}
