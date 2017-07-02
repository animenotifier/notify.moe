package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/utils"
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

// InstallFacebookAuth enables Facebook login for the app.
func InstallFacebookAuth(app *aero.Application) {
	config := &oauth2.Config{
		ClientID:     arn.APIKeys.Facebook.ID,
		ClientSecret: arn.APIKeys.Facebook.Secret,
		RedirectURL:  "https://" + app.Config.Domain + "/auth/facebook/callback",
		Scopes: []string{
			"public_profile",
			"email",
		},
		Endpoint: facebook.Endpoint,
	}

	// Auth
	app.Get("/auth/facebook", func(ctx *aero.Context) string {
		state := ctx.Session().ID()
		url := config.AuthCodeURL(state)
		ctx.Redirect(url)
		return ""
	})

	// Auth Callback
	app.Get("/auth/facebook/callback", func(ctx *aero.Context) string {
		if !ctx.HasSession() {
			return ctx.Error(http.StatusUnauthorized, "Facebook login failed", errors.New("Session does not exist"))
		}

		session := ctx.Session()

		if session.ID() != ctx.Query("state") {
			return ctx.Error(http.StatusUnauthorized, "Facebook login failed", errors.New("Incorrect state"))
		}

		// Handle the exchange code to initiate a transport
		token, err := config.Exchange(oauth2.NoContext, ctx.Query("code"))

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Could not obtain OAuth token", err)
		}

		// Construct the OAuth client
		client := config.Client(oauth2.NoContext, token)

		// Fetch user data from Facebook
		resp, err := client.Get("https://graph.facebook.com/me?fields=email,first_name,last_name,gender")

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Failed requesting user data from Facebook", err)
		}

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		// Construct a FacebookUser object
		fbUser := FacebookUser{}
		err = json.Unmarshal(body, &fbUser)

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Failed parsing user data (JSON)", err)
		}

		// Change googlemail.com to gmail.com
		fbUser.Email = strings.Replace(fbUser.Email, "googlemail.com", "gmail.com", 1)

		// Is this an existing user connecting another social account?
		user := utils.GetUser(ctx)

		if user != nil {
			// Add FacebookToUser reference
			err = user.ConnectFacebook(fbUser.ID)

			if err != nil {
				ctx.Error(http.StatusInternalServerError, "Could not connect account to Facebook account", err)
			}

			authLog.Info("Added Facebook ID to existing account", user.ID, user.Nick, ctx.RealIP(), user.Email, user.RealName())

			return ctx.Redirect("/")
		}

		var getErr error

		// Try to find an existing user via the Facebook user ID
		user, getErr = arn.GetUserFromTable("FacebookToUser", fbUser.ID)

		if getErr == nil && user != nil {
			authLog.Info("User logged in via Facebook ID", user.ID, user.Nick, ctx.RealIP(), user.Email, user.RealName())

			user.LastLogin = arn.DateTimeUTC()
			user.Save()

			session.Set("userId", user.ID)
			return ctx.Redirect("/")
		}

		// Try to find an existing user via the associated e-mail address
		user, getErr = arn.GetUserByEmail(fbUser.Email)

		if getErr == nil && user != nil {
			authLog.Info("User logged in via Email", user.ID, user.Nick, ctx.RealIP(), user.Email, user.RealName())

			user.LastLogin = arn.DateTimeUTC()
			user.Save()

			session.Set("userId", user.ID)
			return ctx.Redirect("/")
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
		err = arn.RegisterUser(user)

		if err != nil {
			ctx.Error(http.StatusInternalServerError, "Could not register a new user", err)
		}

		// Connect account to a Facebook account
		err = user.ConnectFacebook(fbUser.ID)

		if err != nil {
			ctx.Error(http.StatusInternalServerError, "Could not connect account to Facebook account", err)
		}

		// Save user object again with updated data
		user.Save()

		// Login
		session.Set("userId", user.ID)

		// Log
		authLog.Info("Registered new user via Facebook", user.ID, user.Nick, ctx.RealIP(), user.Email, user.RealName())

		// Redirect to frontpage
		return ctx.Redirect("/")
	})
}
