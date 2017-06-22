package auth

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

// InstallGoogleAuth enables Google login for the app.
func InstallGoogleAuth(app *aero.Application) {
	config := &oauth2.Config{
		ClientID:     apiKeys.Google.ID,
		ClientSecret: apiKeys.Google.Secret,
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

		// Try to find an existing user by the Google user ID
		user, getErr := arn.GetUserFromTable("GoogleToUser", googleUser.Sub)

		if getErr == nil && user != nil {
			// Add GoogleToUser reference
			user.Accounts.Google.ID = googleUser.Sub
			arn.DB.Set("GoogleToUser", googleUser.Sub, &arn.GoogleToUser{
				ID:     googleUser.Sub,
				UserID: user.ID,
			})

			session.Set("userId", user.ID)
			return ctx.Redirect("/")
		}

		// Try to find an existing user by the associated e-mail address
		user, getErr = arn.GetUserByEmail(googleUser.Email)

		if getErr == nil && user != nil {
			session.Set("userId", user.ID)
			return ctx.Redirect("/")
		}

		user = arn.NewUser()
		user.Nick = "g" + googleUser.Sub
		user.Email = googleUser.Email
		user.FirstName = googleUser.GivenName
		user.LastName = googleUser.FamilyName
		user.Gender = googleUser.Gender
		user.Accounts.Google.ID = googleUser.Sub
		user.LastLogin = arn.DateTimeUTC()

		arn.PrettyPrint(user)
		// arn.RegisterUser(user)

		return ctx.Error(http.StatusForbidden, "Account does not exist", nil)
	})
}
