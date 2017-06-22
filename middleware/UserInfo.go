package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	println(ctx.GetRequestHeader("Accept-Type"))
	if strings.Index(ctx.GetRequestHeader("Accept-Type"), "text/html") == -1 {
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

		arn.PrettyPrint(user.Browser)

		// OS
		os := parsed.OSInfo()
		user.OS.Name = os.Name
		user.OS.Version = os.Version

		arn.PrettyPrint(user.OS)
	}

	if user.IP != newIP {
		user.IP = newIP
		locationAPI := "https://api.ipinfodb.com/v3/ip-city/?key=" + apiKeys.IPInfoDB.ID + "&ip=" + "2a02:8108:8dc0:3000:6cf1:af03:ce6e:679a" + "&format=json"

		response, data, err := gorequest.New().Get(locationAPI).EndBytes()

		if len(err) > 0 && err[0] != nil {
			color.Red(err[0].Error())
			return
		}

		if response.StatusCode != http.StatusOK {
			body, _ := ioutil.ReadAll(response.Body)
			fmt.Println(response.StatusCode, locationAPI)
			fmt.Println(string(body))
			return
		}

		json.Unmarshal(data, &user.Location)
		arn.PrettyPrint(user.Location)
	}

	// user.Save()
}
