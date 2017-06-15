package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// GoogleUser is the user data we receive from Google
type GoogleUser struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}

// EnableGoogleLogin enables Google login for the app.
func EnableGoogleLogin(app *aero.Application) {
	conf := &oauth2.Config{
		ClientID:     apiKeys.Google.ID,
		ClientSecret: apiKeys.Google.Secret,
		RedirectURL:  "https://beta.notify.moe/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	// Auth
	app.Get("/auth/google", func(ctx *aero.Context) string {
		sessionID := ctx.Session().ID()
		url := conf.AuthCodeURL(sessionID)
		ctx.Redirect(url)
		return ""
	})

	// Auth Callback
	app.Get("/auth/google/callback", func(ctx *aero.Context) string {
		session := ctx.Session()

		if session.ID() != ctx.Query("state") {
			return ctx.Error(http.StatusUnauthorized, "Authorization not allowed for this session", errors.New("Google login failed: Incorrect state"))
		}

		// Handle the exchange code to initiate a transport
		token, err := conf.Exchange(oauth2.NoContext, ctx.Query("code"))

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Could not obtain OAuth token", err)
		}

		// Construct the OAuth client
		client := conf.Client(oauth2.NoContext, token)

		// Fetch user data from Google
		resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Failed requesting user data from Google", err)
		}

		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)

		// Construct a GoogleUser object
		var googleUser GoogleUser
		err = json.Unmarshal(data, &googleUser)

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Failed parsing user data (JSON)", err)
		}

		// Try to find an existing user by the associated e-mail address
		email := googleUser.Email
		user, getErr := arn.GetUserByEmail(email)

		if getErr != nil {
			return ctx.Error(http.StatusForbidden, "Email not registered", err)
		}

		// Login
		session.Set("userId", user.ID)

		// Redirect back to frontpage
		return ctx.Redirect("/")
	})
}
