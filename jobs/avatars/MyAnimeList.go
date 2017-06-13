package main

import (
	"regexp"
	"time"

	"github.com/animenotifier/arn"
	"github.com/parnurzeal/gorequest"
)

var userIDRegex = regexp.MustCompile(`<user_id>(\d+)<\/user_id>`)

// MyAnimeList - https://myanimelist.net/
type MyAnimeList struct {
	RequestLimiter *time.Ticker
}

// GetAvatar returns the Gravatar image for a user (if available).
func (source *MyAnimeList) GetAvatar(user *arn.User) *Avatar {
	malNick := user.Accounts.MyAnimeList.Nick

	// If the user has no username we can't get an avatar.
	if malNick == "" {
		return nil
	}

	// Download user info
	userInfoURL := "https://myanimelist.net/malappinfo.php?u=" + malNick
	response, xml, networkErr := gorequest.New().Get(userInfoURL).End()

	if networkErr != nil || response.StatusCode != 200 {
		return nil
	}

	// Build URL
	matches := userIDRegex.FindStringSubmatch(xml)

	if matches == nil || len(matches) < 2 {
		return nil
	}

	malID := matches[1]
	malAvatarURL := "https://myanimelist.cdn-dena.com/images/userimages/" + malID + ".jpg"

	// Wait for request limiter to allow us to send a request
	<-source.RequestLimiter.C

	// Download
	return AvatarFromURL(malAvatarURL, user)
}
