package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aerogo/aero"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// EnableGoogleLogin enables Google login for the app.
func EnableGoogleLogin(app *aero.Application) {
	var api APIKeys
	data, _ := ioutil.ReadFile("security/api-keys.json")
	json.Unmarshal(data, &api)

	conf := &oauth2.Config{
		ClientID:     api.Google.ID,
		ClientSecret: api.Google.Secret,
		RedirectURL:  "https://beta.notify.moe/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	// Auth
	app.Get("/auth/google", func(ctx *aero.Context) string {
		url := conf.AuthCodeURL(ctx.Session().ID())
		ctx.Redirect(url)
		return ""
	})

	// Auth Callback
	app.Get("/auth/google/callback", func(ctx *aero.Context) string {
		if ctx.Session().ID() != ctx.Query("state") {
			return ctx.Error(http.StatusBadRequest, "Authorization not allowed for this session", errors.New("Google login failed: Incorrect state"))
		}

		// Handle the exchange code to initiate a transport
		token, err := conf.Exchange(oauth2.NoContext, ctx.Query("code"))
		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Could not obtain OAuth token", err)
		}

		// Construct the OAuth client
		client := conf.Client(oauth2.NoContext, token)

		resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Failed requesting user data from Google", err)
		}
		defer resp.Body.Close()
		dataBytes, _ := ioutil.ReadAll(resp.Body)
		data := string(dataBytes)
		log.Println("Resp body: ", data)

		return ctx.Text(data)
	})
}
