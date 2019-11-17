package auth

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aerogo/aero"
	"github.com/aerogo/log"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// GoogleUser is the user data we receive from Google
type GoogleUser struct {
	Sub        string `json:"sub"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`

	// Name          string `json:"name"`
	// Profile       string `json:"profile"`
	// Picture       string `json:"picture"`
	// EmailVerified bool   `json:"email_verified"`
}

// Google enables Google login for the app.
func Google(app *aero.Application, authLog *log.Log) {
	// OAuth2 configuration defines the API keys,
	// scopes of required data and the redirect URL
	// that Google should send the user to after
	// a successful login on their pages.
	config := &oauth2.Config{
		ClientID:     arn.APIKeys.Google.ID,
		ClientSecret: arn.APIKeys.Google.Secret,
		RedirectURL:  "https://" + assets.Domain + "/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			// "https://www.googleapis.com/auth/plus.me",
			// "https://www.googleapis.com/auth/plus.login",
		},
		Endpoint: google.Endpoint,
	}

	// When a user visits /auth/google, we ask OAuth2 config for a URL
	// to redirect the user to. Once the user has logged in on that page,
	// he'll be redirected back to our servers to the callback page.
	app.Get("/auth/google", func(ctx aero.Context) error {
		state := ctx.Session().ID()
		url := config.AuthCodeURL(state)
		return ctx.Redirect(http.StatusTemporaryRedirect, url)
	})

	// This is the redirect URL that we specified in the OAuth2 config.
	// The user has successfully completed the login on Google servers.
	// Now we have to check for fraud requests and request user information.
	// If both Google ID and email can't be found in our DB, register a new user.
	// Otherwise, log in the user with the given Google ID or email.
	app.Get("/auth/google/callback", func(ctx aero.Context) error {
		if !ctx.HasSession() {
			return ctx.Error(http.StatusUnauthorized, "Google login failed", errors.New("Session does not exist"))
		}

		session := ctx.Session()

		if session.ID() != ctx.Query("state") {
			return ctx.Error(http.StatusUnauthorized, "Google login failed", errors.New("Incorrect state"))
		}

		// Handle the exchange code to initiate a transport
		token, err := config.Exchange(context.Background(), ctx.Query("code"))

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Could not obtain OAuth token", err)
		}

		// Construct the OAuth client
		client := config.Client(context.Background(), token)

		// Fetch user data from Google
		response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Failed requesting user data from Google", err)
		}

		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)

		// Construct a GoogleUser object
		var googleUser GoogleUser
		err = jsoniter.Unmarshal(data, &googleUser)

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Failed parsing user data (JSON)", err)
		}

		if googleUser.Sub == "" {
			return ctx.Error(http.StatusBadRequest, "Failed retrieving Google data", errors.New("Empty ID"))
		}

		// Change googlemail.com to gmail.com
		googleUser.Email = strings.Replace(googleUser.Email, "googlemail.com", "gmail.com", 1)

		// Is this an existing user connecting another social account?
		user := arn.GetUserFromContext(ctx)

		if user != nil {
			// Add GoogleToUser reference
			user.ConnectGoogle(googleUser.Sub)

			// Save in DB
			user.Save()

			// Log
			authLog.Info("Added Google ID to existing account | %s | %s | %s | %s | %s", user.Nick, user.ID, ctx.IP(), user.Email, user.RealName())

			return ctx.Redirect(http.StatusTemporaryRedirect, "/")
		}

		var getErr error

		// Try to find an existing user via the Google user ID
		user, getErr = arn.GetUserByGoogleID(googleUser.Sub)

		if getErr == nil && user != nil {
			authLog.Info("User logged in via Google ID | %s | %s | %s | %s | %s", user.Nick, user.ID, ctx.IP(), user.Email, user.RealName())

			user.LastLogin = arn.DateTimeUTC()
			user.Save()

			session.Set("userId", user.ID)
			return ctx.Redirect(http.StatusTemporaryRedirect, "/")
		}

		// Try to find an existing user via the associated e-mail address
		user, getErr = arn.GetUserByEmail(googleUser.Email)

		if getErr == nil && user != nil {
			authLog.Info("User logged in via Email | %s | %s | %s | %s | %s", user.Nick, user.ID, ctx.IP(), user.Email, user.RealName())

			// Add GoogleToUser reference
			user.ConnectGoogle(googleUser.Sub)

			user.LastLogin = arn.DateTimeUTC()
			user.Save()

			session.Set("userId", user.ID)
			return ctx.Redirect(http.StatusTemporaryRedirect, "/")
		}

		// Register new user
		user = arn.NewUser()
		user.Nick = "g" + googleUser.Sub
		user.Email = googleUser.Email
		user.FirstName = googleUser.GivenName
		user.LastName = googleUser.FamilyName
		user.Gender = googleUser.Gender
		user.LastLogin = arn.DateTimeUTC()

		// Save basic user info already to avoid data inconsistency problems
		user.Save()

		// Register user
		arn.RegisterUser(user)

		// Connect account to a Google account
		user.ConnectGoogle(googleUser.Sub)

		// Save user object again with updated data
		user.Save()

		// Login
		session.Set("userId", user.ID)

		// Log
		authLog.Info("Registered new user via Google | %s | %s | %s | %s | %s", user.Nick, user.ID, ctx.IP(), user.Email, user.RealName())

		// Redirect to starting page for new users
		return ctx.Redirect(http.StatusTemporaryRedirect, newUserStartRoute)
	})
}
