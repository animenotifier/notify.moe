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
	"golang.org/x/oauth2/facebook"
)

// FacebookUser is the user data we receive from Facebook
type FacebookUser struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
}

// Facebook enables Facebook login for the app.
func Facebook(app *aero.Application, authLog *log.Log) {
	// OAuth2 configuration defines the API keys,
	// scopes of required data and the redirect URL
	// that Facebook should send the user to after
	// a successful login on their pages.
	config := &oauth2.Config{
		ClientID:     arn.APIKeys.Facebook.ID,
		ClientSecret: arn.APIKeys.Facebook.Secret,
		RedirectURL:  "https://" + assets.Domain + "/auth/facebook/callback",
		Scopes: []string{
			"public_profile",
			"email",
		},
		Endpoint: facebook.Endpoint,
	}

	// When a user visits /auth/facebook, we ask OAuth2 config for a URL
	// to redirect the user to. Once the user has logged in on that page,
	// he'll be redirected back to our servers to the callback page.
	app.Get("/auth/facebook", func(ctx aero.Context) error {
		state := ctx.Session().ID()
		url := config.AuthCodeURL(state)
		return ctx.Redirect(http.StatusTemporaryRedirect, url)
	})

	// This is the redirect URL that we specified in the OAuth2 config.
	// The user has successfully completed the login on Facebook servers.
	// Now we have to check for fraud requests and request user information.
	// If both Facebook ID and email can't be found in our DB, register a new user.
	// Otherwise, log in the user with the given Facebook ID or email.
	app.Get("/auth/facebook/callback", func(ctx aero.Context) error {
		if !ctx.HasSession() {
			return ctx.Error(http.StatusUnauthorized, "Facebook login failed", errors.New("Session does not exist"))
		}

		session := ctx.Session()

		if session.ID() != ctx.Query("state") {
			return ctx.Error(http.StatusUnauthorized, "Facebook login failed", errors.New("Incorrect state"))
		}

		// Handle the exchange code to initiate a transport
		token, err := config.Exchange(context.Background(), ctx.Query("code"))

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Could not obtain OAuth token", err)
		}

		// Construct the OAuth client
		client := config.Client(context.Background(), token)

		// Fetch user data from Facebook
		response, err := client.Get("https://graph.facebook.com/me?fields=email,first_name,last_name,gender")

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Failed requesting user data from Facebook", err)
		}

		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)

		// Construct a FacebookUser object
		fbUser := FacebookUser{}
		err = jsoniter.Unmarshal(body, &fbUser)

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Failed parsing user data (JSON)", err)
		}

		// Change googlemail.com to gmail.com
		fbUser.Email = strings.Replace(fbUser.Email, "googlemail.com", "gmail.com", 1)

		// Is this an existing user connecting another social account?
		user := arn.GetUserFromContext(ctx)

		if user != nil {
			// Add FacebookToUser reference
			user.ConnectFacebook(fbUser.ID)

			// Save in DB
			user.Save()

			// Log
			authLog.Info("Added Facebook ID to existing account | %s | %s | %s | %s | %s", user.Nick, user.ID, ctx.IP(), user.Email, user.RealName())

			return ctx.Redirect(http.StatusTemporaryRedirect, "/")
		}

		var getErr error

		// Try to find an existing user via the Facebook user ID
		user, getErr = arn.GetUserByFacebookID(fbUser.ID)

		if getErr == nil && user != nil {
			authLog.Info("User logged in via Facebook ID | %s | %s | %s | %s | %s", user.Nick, user.ID, ctx.IP(), user.Email, user.RealName())

			// Add FacebookToUser reference
			user.ConnectFacebook(fbUser.ID)

			user.LastLogin = arn.DateTimeUTC()
			user.Save()

			session.Set("userId", user.ID)
			return ctx.Redirect(http.StatusTemporaryRedirect, "/")
		}

		// Try to find an existing user via the associated e-mail address
		user, getErr = arn.GetUserByEmail(fbUser.Email)

		if getErr == nil && user != nil {
			authLog.Info("User logged in via Email | %s | %s | %s | %s | %s", user.Nick, user.ID, ctx.IP(), user.Email, user.RealName())

			user.LastLogin = arn.DateTimeUTC()
			user.Save()

			session.Set("userId", user.ID)
			return ctx.Redirect(http.StatusTemporaryRedirect, "/")
		}

		// Register new user
		user = arn.NewUser()
		user.Nick = "fb" + fbUser.ID
		user.Email = fbUser.Email
		user.FirstName = fbUser.FirstName
		user.LastName = fbUser.LastName
		user.Gender = fbUser.Gender
		user.LastLogin = arn.DateTimeUTC()

		// Save basic user info already to avoid data inconsistency problems
		user.Save()

		// Register user
		arn.RegisterUser(user)

		// Connect account to a Facebook account
		user.ConnectFacebook(fbUser.ID)

		// Save user object again with updated data
		user.Save()

		// Login
		session.Set("userId", user.ID)

		// Log
		authLog.Info("Registered new user via Facebook | %s | %s | %s | %s | %s", user.Nick, user.ID, ctx.IP(), user.Email, user.RealName())

		// Redirect to starting page for new users
		return ctx.Redirect(http.StatusTemporaryRedirect, newUserStartRoute)
	})
}
