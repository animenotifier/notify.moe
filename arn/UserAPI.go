package arn

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/animenotifier/notify.moe/arn/validate"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
	"github.com/aerogo/http/client"
	"github.com/animenotifier/notify.moe/arn/autocorrect"
)

// Force interface implementations
var (
	_ PostParent   = (*User)(nil)
	_ api.Editable = (*User)(nil)
)

// Authorize returns an error if the given API POST request is not authorized.
func (user *User) Authorize(ctx aero.Context, action string) error {
	editor := GetUserFromContext(ctx)

	if editor == nil {
		return errors.New("Not authorized")
	}

	if editor.ID != ctx.Get("id") && editor.Role != "admin" {
		return errors.New("Can not modify data from other users")
	}

	return nil
}

// Edit updates the user object.
func (user *User) Edit(ctx aero.Context, key string, value reflect.Value, newValue reflect.Value) (bool, error) {
	switch key {
	case "Nick":
		newNick := newValue.String()
		err := user.SetNick(newNick)
		return true, err

	case "Email":
		newEmail := newValue.String()
		err := user.SetEmail(newEmail)
		return true, err

	case "Gender":
		newGender := newValue.String()

		if newGender != "male" && newGender != "female" {
			return true, errors.New("Invalid gender")
		}

		user.Gender = newGender
		return true, nil

	case "Website":
		newSite := newValue.String()

		if newSite == "" {
			user.Website = newSite
			return true, nil
		}

		if autocorrect.IsTrackerLink(newSite) {
			return true, errors.New("Not an actual personal website or homepage")
		}

		newSite = autocorrect.Website(newSite)

		if !validate.URI("https://" + newSite) {
			return true, errors.New("Not a valid website link")
		}

		response, err := client.Get("https://" + newSite).End()

		if err != nil || response.StatusCode() >= 400 {
			return true, fmt.Errorf("https://%s seems to be inaccessible", newSite)
		}

		user.Website = newSite
		return true, nil

	case "BirthDay":
		newBirthDay := newValue.String()

		if AgeInYears(newBirthDay) <= 0 {
			return true, errors.New("Invalid birthday (make sure to use YYYY-MM-DD format, e.g. 2000-01-17)")
		}

		user.BirthDay = newBirthDay
		return true, nil

	case "ProExpires", "Avatar.Extension", "Avatar.LastModified":
		editor := GetUserFromContext(ctx)

		if editor == nil || editor.Role != "admin" {
			return true, errors.New("Not authorized to edit")
		}

	case "Accounts.Discord.Nick":
		newNick := newValue.String()

		if newNick == "" {
			value.SetString(newNick)
			user.Accounts.Discord.Verified = false
			return true, nil
		}

		if !validate.DiscordNick(newNick) {
			return true, errors.New("Discord username must include your name and the 4-digit Discord tag (e.g. Yandere#1234)")
		}

		// Trim spaces
		parts := strings.Split(newNick, "#")
		parts[0] = strings.TrimSpace(parts[0])
		parts[1] = strings.TrimSpace(parts[1])
		newNick = strings.Join(parts, "#")

		if value.String() != newNick {
			value.SetString(newNick)
			user.Accounts.Discord.Verified = false
		}

		return true, nil
	}

	// Automatically correct account nicks
	if strings.HasPrefix(key, "Accounts.") && strings.HasSuffix(key, ".Nick") {
		newNick := newValue.String()
		newNick = autocorrect.AccountNick(newNick)
		value.SetString(newNick)
		return true, nil
	}

	return false, nil
}

// Save saves the user object in the database.
func (user *User) Save() {
	DB.Set("User", user.ID, user)
}

// Filter removes privacy critical fields from the user object.
func (user *User) Filter() {
	user.Email = ""
	user.Gender = ""
	user.FirstName = ""
	user.LastName = ""
	user.IP = ""
	user.UserAgent = ""
	user.LastLogin = ""
	user.Accounts.Facebook.ID = ""
	user.Accounts.Google.ID = ""
	user.Accounts.Twitter.ID = ""
	user.BirthDay = ""
	user.Location = &Location{}
	user.Browser = UserBrowser{}
	user.OS = UserOS{}
}

// ShouldFilter tells whether data needs to be filtered in the given context.
func (user *User) ShouldFilter(ctx aero.Context) bool {
	ctxUser := GetUserFromContext(ctx)

	if ctxUser != nil && ctxUser.Role == "admin" {
		return false
	}

	return true
}
