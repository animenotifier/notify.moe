package auth

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/aerogo/aero"
	"github.com/aerogo/log"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
	"github.com/gomodule/oauth1/oauth"
	jsoniter "github.com/json-iterator/go"
)

// TwitterUser is the user data we receive from Twitter
type TwitterUser struct {
	ID          string `json:"id_str"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ScreenName  string `json:"screen_name"`
}

// Twitter enables Twitter login for the app.
func Twitter(app *aero.Application, authLog *log.Log) {
	// oauth1 configuration defines the API keys,
	// the url for the request token, the access token and the authorisation.
	config := &oauth.Client{
		TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
		ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authenticate",
		TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
		Credentials: oauth.Credentials{
			Token:  arn.APIKeys.Twitter.ID,
			Secret: arn.APIKeys.Twitter.Secret,
		},
	}

	// When a user visits /auth/twitter, we ask OAuth1 to give us
	// a request token and give us a URL to redirect the user to.
	// Once the user has approved the application on that page,
	// he'll be redirected back to our servers to the callback page.
	app.Get("/auth/twitter", func(ctx aero.Context) error {
		callback := "https://" + assets.Domain + "/auth/twitter/callback"
		tempCred, err := config.RequestTemporaryCredentials(nil, callback, nil)

		if err != nil {
			fmt.Printf("Error: %s", err)
		}

		ctx.Session().Set("tempCred", tempCred)
		url := config.AuthorizationURL(tempCred, nil)
		return ctx.Redirect(http.StatusTemporaryRedirect, url)
	})

	// This is the redirect URL that we specified in /auth/twitter.
	// The user has allowed the application to have access to his data.
	// Now we have to check for fraud requests and request user information.
	// If both Twitter ID and email can't be found in our DB, register a new user.
	// Otherwise, log in the user with the given Twitter ID or email.
	app.Get("/auth/twitter/callback", func(ctx aero.Context) error {
		if !ctx.HasSession() {
			return ctx.Error(http.StatusUnauthorized, "Twitter login failed", errors.New("Session does not exist"))
		}

		session := ctx.Session()

		// Get back the request token to get the access token
		tempCred, _ := session.Get("tempCred").(*oauth.Credentials)
		session.Delete("tempCred")

		if tempCred == nil || tempCred.Token != ctx.Query("oauth_token") {
			return ctx.Error(http.StatusBadRequest, "Unknown OAuth request token", nil)
		}

		// Get the request token from twitter
		tokenCred, _, err := config.RequestToken(nil, tempCred, ctx.Query("oauth_verifier"))

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Could not obtain OAuth access token", err)
		}

		// Fetch user data from Twitter
		params := url.Values{
			"include_email": {"true"},
			"skip_status":   {"true"},
		}

		response, err := config.Get(nil, tokenCred, "https://api.twitter.com/1.1/account/verify_credentials.json", params)

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Failed requesting user data from Twitter", err)
		}

		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))

		// Construct a TwitterUser object
		twUser := TwitterUser{}
		err = jsoniter.Unmarshal(body, &twUser)

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Failed parsing user data (JSON)", err)
		}

		// Change googlemail.com to gmail.com
		twUser.Email = strings.Replace(twUser.Email, "googlemail.com", "gmail.com", 1)

		// Is this an existing user connecting another social account?
		user := arn.GetUserFromContext(ctx)

		if user != nil {
			// Add TwitterToUser reference
			user.ConnectTwitter(twUser.ID)

			// Save in DB
			user.Accounts.Twitter.Nick = twUser.ScreenName
			user.Save()

			// Log
			authLog.Info("Added Twitter ID to existing account | %s | %s | %s | %s | %s", user.Nick, user.ID, ctx.IP(), user.Email, user.RealName())

			return ctx.Redirect(http.StatusTemporaryRedirect, "/")
		}

		var getErr error

		// Try to find an existing user via the Twitter user ID
		user, getErr = arn.GetUserByTwitterID(twUser.ID)

		if getErr == nil && user != nil {
			authLog.Info("User logged in via Twitter ID | %s | %s | %s | %s | %s", user.Nick, user.ID, ctx.IP(), user.Email, user.RealName())

			user.LastLogin = arn.DateTimeUTC()
			user.Save()

			session.Set("userId", user.ID)
			return ctx.Redirect(http.StatusTemporaryRedirect, "/")
		}

		// Try to find an existing user via the associated e-mail address
		user, getErr = arn.GetUserByEmail(twUser.Email)

		if getErr == nil && user != nil {
			authLog.Info("User logged in via Email | %s | %s | %s | %s | %s", user.Nick, user.ID, ctx.IP(), user.Email, user.RealName())

			// Add TwitterToUser reference
			user.ConnectTwitter(twUser.ID)

			user.Accounts.Twitter.Nick = twUser.ScreenName
			user.LastLogin = arn.DateTimeUTC()
			user.Save()

			session.Set("userId", user.ID)
			return ctx.Redirect(http.StatusTemporaryRedirect, "/")
		}

		// Register new user
		user = arn.NewUser()
		user.Nick = "tw" + twUser.ID
		user.Email = twUser.Email
		user.LastLogin = arn.DateTimeUTC()
		user.FirstName = twUser.Name

		// Save basic user info already to avoid data inconsistency problems
		user.Save()

		// Register user
		arn.RegisterUser(user)

		// Connect account to a Twitter account
		user.ConnectTwitter(twUser.ID)

		// Copy fields
		user.Accounts.Twitter.Nick = twUser.ScreenName
		user.Introduction = twUser.Description

		// Save user object again with updated data
		user.Save()

		// Login
		session.Set("userId", user.ID)

		// Log
		authLog.Info("Registered new user via Twitter | %s | %s | %s | %s | %s", user.Nick, user.ID, ctx.IP(), user.Email, user.RealName())

		// Redirect to starting page for new users
		return ctx.Redirect(http.StatusTemporaryRedirect, newUserStartRoute)
	})
}
