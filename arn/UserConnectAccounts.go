package arn

// ConnectGoogle connects the user's account with a Google account.
func (user *User) ConnectGoogle(googleID string) {
	if googleID == "" {
		return
	}

	user.Accounts.Google.ID = googleID

	DB.Set("GoogleToUser", googleID, &GoogleToUser{
		ID:     googleID,
		UserID: user.ID,
	})
}

// ConnectFacebook connects the user's account with a Facebook account.
func (user *User) ConnectFacebook(facebookID string) {
	if facebookID == "" {
		return
	}

	user.Accounts.Facebook.ID = facebookID

	DB.Set("FacebookToUser", facebookID, &FacebookToUser{
		ID:     facebookID,
		UserID: user.ID,
	})
}

// ConnectTwitter connects the user's account with a Twitter account.
func (user *User) ConnectTwitter(twtterID string) {
	if twtterID == "" {
		return
	}

	user.Accounts.Twitter.ID = twtterID

	DB.Set("TwitterToUser", twtterID, &TwitterToUser{
		ID:     twtterID,
		UserID: user.ID,
	})
}
