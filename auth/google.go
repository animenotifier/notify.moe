package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/utils"
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

// InstallGoogleAuth enables Google login for the app.
func InstallGoogleAuth(app *aero.Application) {
	config := &oauth2.Config{
		ClientID:     arn.APIKeys.Google.ID,
		ClientSecret: arn.APIKeys.Google.Secret,
		RedirectURL:  "https://" + app.Config.Domain + "/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			// "https://www.googleapis.com/auth/plus.me",
			// "https://www.googleapis.com/auth/plus.login",
		},
		Endpoint: google.Endpoint,
	}

	// Auth
	app.Get("/auth/google", func(ctx *aero.Context) string {
		sessionID := ctx.Session().ID()
		url := config.AuthCodeURL(sessionID)
		ctx.Redirect(url)
		return ""
	})

	// Auth Callback
	app.Get("/auth/google/callback", func(ctx *aero.Context) string {
		if !ctx.HasSession() {
			return ctx.Error(http.StatusUnauthorized, "Google login failed", errors.New("Session does not exist"))
		}

		session := ctx.Session()

		if session.ID() != ctx.Query("state") {
			return ctx.Error(http.StatusUnauthorized, "Google login failed", errors.New("Incorrect state"))
		}

		// Handle the exchange code to initiate a transport
		token, err := config.Exchange(oauth2.NoContext, ctx.Query("code"))

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Could not obtain OAuth token", err)
		}

		// Construct the OAuth client
		client := config.Client(oauth2.NoContext, token)

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

		// Is this an existing user connecting another social account?
		user := utils.GetUser(ctx)

		if user != nil {
			println("Connected")

			// Add GoogleToUser reference
			err = user.ConnectGoogle(googleUser.Sub)

			if err != nil {
				ctx.Error(http.StatusInternalServerError, "Could not connect account to Google account", err)
			}

			return ctx.Redirect("/")
		}

		var getErr error

		// Try to find an existing user via the Google user ID
		user, getErr = arn.GetUserFromTable("GoogleToUser", googleUser.Sub)

		if getErr == nil && user != nil {
			authLog.Info("User logged in via Google ID", user.ID, user.Nick, ctx.RealIP(), user.Email, user.RealName())

			user.LastLogin = arn.DateTimeUTC()
			user.Save()

			session.Set("userId", user.ID)
			return ctx.Redirect("/")
		}

		// Try to find an existing user via the associated e-mail address
		user, getErr = arn.GetUserByEmail(googleUser.Email)

		if getErr == nil && user != nil {
			authLog.Info("User logged in via Email", user.ID, user.Nick, ctx.RealIP(), user.Email, user.RealName())

			user.LastLogin = arn.DateTimeUTC()
			user.Save()

			session.Set("userId", user.ID)
			return ctx.Redirect("/")
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
		err = arn.RegisterUser(user)

		if err != nil {
			ctx.Error(http.StatusInternalServerError, "Could not register a new user", err)
		}

		// Connect account to a Google account
		err = user.ConnectGoogle(googleUser.Sub)

		if err != nil {
			ctx.Error(http.StatusInternalServerError, "Could not connect account to Google account", err)
		}

		// Save user object again with updated data
		user.Save()

		// Login
		session.Set("userId", user.ID)

		// Log
		authLog.Info("Registered new user", user.ID, user.Nick, ctx.RealIP(), user.Email, user.RealName())

		// Redirect to frontpage
		return ctx.Redirect("/")
	})
}
