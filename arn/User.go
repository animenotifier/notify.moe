package arn

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/aerogo/aero/event"
	"github.com/aerogo/http/client"
	"github.com/animenotifier/notify.moe/arn/autocorrect"
	"github.com/animenotifier/notify.moe/arn/validate"
	gravatar "github.com/ungerik/go-gravatar"
)

var (
	setNickMutex  sync.Mutex
	setEmailMutex sync.Mutex
)

// UserID represents a user ID.
type UserID = ID

// User is a registered person.
type User struct {
	ID           UserID       `json:"id" primary:"true"`
	Nick         string       `json:"nick" editable:"true"`
	FirstName    string       `json:"firstName" private:"true"`
	LastName     string       `json:"lastName" private:"true"`
	Email        string       `json:"email" editable:"true" private:"true"`
	Role         string       `json:"role"`
	Registered   string       `json:"registered"`
	LastLogin    string       `json:"lastLogin" private:"true"`
	LastSeen     string       `json:"lastSeen" private:"true"`
	ProExpires   string       `json:"proExpires" editable:"true"`
	Gender       string       `json:"gender" editable:"true" private:"true" datalist:"genders"`
	Language     string       `json:"language"`
	Introduction string       `json:"introduction" editable:"true" type:"textarea"`
	Website      string       `json:"website" editable:"true"`
	BirthDay     string       `json:"birthDay" editable:"true" private:"true"`
	IP           string       `json:"ip" private:"true"`
	UserAgent    string       `json:"agent" private:"true"`
	Balance      int          `json:"balance" private:"true"`
	Image        Image        `json:"image"`
	Avatar       UserAvatar   `json:"avatar" editable:"true"`
	Cover        UserCover    `json:"cover"`
	Accounts     UserAccounts `json:"accounts" private:"true"`
	Browser      UserBrowser  `json:"browser" private:"true"`
	OS           UserOS       `json:"os" private:"true"`
	Location     *Location    `json:"location" private:"true"`
	FollowIDs    []UserID     `json:"follows"`

	hasPosts

	eventStreams struct {
		sync.Mutex
		value []*event.Stream
	}
}

// NewUser creates an empty user object with a unique ID.
func NewUser() *User {
	user := &User{
		ID: GenerateID("User"),

		// Avoid nil value fields
		Location: &Location{},
	}

	return user
}

// RegisterUser registers a new user in the database and sets up all the required references.
func RegisterUser(user *User) {
	user.Registered = DateTimeUTC()
	user.LastLogin = user.Registered
	user.LastSeen = user.Registered

	// Save nick in NickToUser collection
	DB.Set("NickToUser", user.Nick, &NickToUser{
		Nick:   user.Nick,
		UserID: user.ID,
	})

	// Save email in EmailToUser collection
	if user.Email != "" {
		DB.Set("EmailToUser", user.Email, &EmailToUser{
			Email:  user.Email,
			UserID: user.ID,
		})
	}

	// Create default settings
	NewSettings(user).Save()

	// Add empty anime list
	DB.Set("AnimeList", user.ID, &AnimeList{
		UserID: user.ID,
		Items:  []*AnimeListItem{},
	})

	// Add empty inventory
	NewInventory(user.ID).Save()

	// Add draft index
	NewDraftIndex(user.ID).Save()

	// Add empty push subscriptions
	DB.Set("PushSubscriptions", user.ID, &PushSubscriptions{
		UserID: user.ID,
		Items:  []*PushSubscription{},
	})

	// Add empty notifications list
	NewUserNotifications(user.ID).Save()

	// Fetch gravatar
	if user.Email != "" && !IsDevelopment() {
		gravatarURL := gravatar.Url(user.Email) + "?s=" + fmt.Sprint(AvatarMaxSize) + "&d=404&r=pg"
		gravatarURL = strings.Replace(gravatarURL, "http://", "https://", 1)

		response, err := client.Get(gravatarURL).End()

		if err == nil && response.StatusCode() == http.StatusOK {
			data := response.Bytes()
			err = user.SetImageBytes(data)

			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// SendNotification accepts a PushNotification and generates a new Notification object.
// The notification is then sent to all registered push devices.
func (user *User) SendNotification(pushNotification *PushNotification) {
	// Don't ever send notifications in development mode
	if IsDevelopment() && user.ID != "4J6qpK1ve" {
		return
	}

	// Save notification in database
	notification := NewNotification(user.ID, pushNotification)
	notification.Save()

	userNotifications := user.Notifications()
	err := userNotifications.Add(notification.ID)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNotifications.Save()

	// Send push notification
	subs := user.PushSubscriptions()
	expired := []*PushSubscription{}

	for _, sub := range subs.Items {
		if sub.Endpoint == "" {
			expired = append(expired, sub)
			continue
		}

		response, err := sub.SendNotification(pushNotification)

		// It is possible to receive a non-nil response with an error, so check the status.
		if response != nil && (response.StatusCode == http.StatusGone || response.StatusCode == http.StatusForbidden) {
			expired = append(expired, sub)
		}

		// Print errors
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Response body must be closed if no error was returned
		defer response.Body.Close()

		// Print bad status codes
		if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
			body, _ := ioutil.ReadAll(response.Body)
			fmt.Println(response.StatusCode, string(body))
			continue
		}

		sub.LastSuccess = DateTimeUTC()
	}

	// Remove expired items
	if len(expired) > 0 {
		for _, sub := range expired {
			subs.Remove(sub.ID())
		}
	}

	// Save changes
	subs.Save()

	// Send an event to the user's open tabs
	notificationCount := event.New("notificationCount", userNotifications.CountUnseen())
	user.BroadcastEvent(notificationCount)
}

// RealName returns the real name of the user.
func (user *User) RealName() string {
	if user.LastName == "" {
		return user.FirstName
	}

	if user.FirstName == "" {
		return user.LastName
	}

	return user.FirstName + " " + user.LastName
}

// RegisteredTime returns the time the user registered his account.
func (user *User) RegisteredTime() time.Time {
	reg, _ := time.Parse(time.RFC3339, user.Registered)
	return reg
}

// LastSeenTime returns the time the user was last seen on the site.
func (user *User) LastSeenTime() time.Time {
	lastSeen, _ := time.Parse(time.RFC3339, user.LastSeen)
	return lastSeen
}

// AgeInYears returns the user's age in years.
func (user *User) AgeInYears() int {
	return AgeInYears(user.BirthDay)
}

// IsActive tells you whether the user is active.
func (user *User) IsActive() bool {
	lastSeen, _ := time.Parse(time.RFC3339, user.LastSeen)
	twoWeeksAgo := time.Now().Add(-14 * 24 * time.Hour)

	if lastSeen.Unix() < twoWeeksAgo.Unix() {
		return false
	}

	if len(user.AnimeList().Items) == 0 {
		return false
	}

	return true
}

// IsPro returns whether the user is a PRO user or not.
func (user *User) IsPro() bool {
	if user.ProExpires == "" {
		return false
	}

	return DateTimeUTC() < user.ProExpires
}

// ExtendProDuration extends the PRO account duration by the given duration.
func (user *User) ExtendProDuration(duration time.Duration) {
	now := time.Now().UTC()
	expires, _ := time.Parse(time.RFC3339, user.ProExpires)

	// If the user never had a PRO account yet,
	// or if it already expired,
	// use the current time as the start time.
	if user.ProExpires == "" || now.Unix() > expires.Unix() {
		expires = now
	}

	user.ProExpires = expires.Add(duration).Format(time.RFC3339)
}

// TimeSinceRegistered returns the duration since the user registered his account.
func (user *User) TimeSinceRegistered() time.Duration {
	registered, _ := time.Parse(time.RFC3339, user.Registered)
	return time.Since(registered)
}

// HasNick returns whether the user has a custom nickname.
func (user *User) HasNick() bool {
	return !strings.HasPrefix(user.Nick, "g") && !strings.HasPrefix(user.Nick, "fb") && !strings.HasPrefix(user.Nick, "t") && user.Nick != ""
}

// WebsiteURL adds https:// to the URL.
func (user *User) WebsiteURL() string {
	return "https://" + user.WebsiteShortURL()
}

// WebsiteShortURL returns the user website without the protocol.
func (user *User) WebsiteShortURL() string {
	return strings.Replace(strings.Replace(user.Website, "https://", "", 1), "http://", "", 1)
}

// Link returns the URI to the user page.
func (user *User) Link() string {
	return "/+" + user.Nick
}

// GetID returns the ID.
func (user *User) GetID() string {
	return user.ID
}

// HasAvatar tells you whether the user has an avatar or not.
func (user *User) HasAvatar() bool {
	return user.Avatar.Extension != ""
}

// AvatarLink returns the URL to the user avatar.
// Expects "small" (50 x 50) or "large" (560 x 560) for the size parameter.
func (user *User) AvatarLink(size string) string {
	if user.HasAvatar() {
		return fmt.Sprintf("//%s/images/avatars/%s/%s%s?%v", MediaHost, size, user.ID, user.Avatar.Extension, user.Avatar.LastModified)
	}

	return fmt.Sprintf("//%s/images/elements/no-avatar.svg", MediaHost)
}

// CoverLink ...
func (user *User) CoverLink(size string) string {
	if user.Cover.Extension != "" && user.IsPro() {
		return fmt.Sprintf("//%s/images/covers/%s/%s%s?%v", MediaHost, size, user.ID, user.Cover.Extension, user.Cover.LastModified)
	}

	return "/images/elements/default-cover.jpg"
}

// Gravatar returns the URL to the gravatar if an email has been registered.
func (user *User) Gravatar() string {
	if user.Email == "" {
		return ""
	}

	return gravatar.SecureUrl(user.Email) + "?s=" + fmt.Sprint(AvatarMaxSize)
}

// EditorScore returns the editor score.
func (user *User) EditorScore() int {
	ignoreDifferences := FilterIgnoreAnimeDifferences(func(entry *IgnoreAnimeDifference) bool {
		return entry.CreatedBy == user.ID
	})

	score := len(ignoreDifferences) * IgnoreAnimeDifferenceEditorScore

	logEntries := FilterEditLogEntries(func(entry *EditLogEntry) bool {
		return entry.UserID == user.ID
	})

	for _, entry := range logEntries {
		score += entry.EditorScore()
	}

	return score
}

// ActivateItemEffect activates an item in the user inventory by the given item ID.
func (user *User) ActivateItemEffect(itemID string) error {
	month := 30 * 24 * time.Hour

	switch itemID {
	case "pro-account-1":
		user.ExtendProDuration(1 * month)
		user.Save()
		return nil

	case "pro-account-3":
		user.ExtendProDuration(3 * month)
		user.Save()
		return nil

	case "pro-account-6":
		user.ExtendProDuration(6 * month)
		user.Save()
		return nil

	case "pro-account-12":
		user.ExtendProDuration(12 * month)
		user.Save()
		return nil

	case "pro-account-24":
		user.ExtendProDuration(24 * month)
		user.Save()
		return nil

	default:
		return errors.New("Can't activate unknown item: " + itemID)
	}
}

// SetNick changes the user's nickname safely.
func (user *User) SetNick(newName string) error {
	setNickMutex.Lock()
	defer setNickMutex.Unlock()

	newName = autocorrect.UserNick(newName)

	if !validate.Nick(newName) {
		return errors.New("Invalid nickname")
	}

	if newName == user.Nick {
		return nil
	}

	// Make sure the nickname doesn't exist already
	_, err := GetUserByNick(newName)

	// If there was no error: the username exists.
	// If "not found" is not included in the error message it's a different error type.
	if err == nil || !strings.Contains(err.Error(), "not found") {
		return errors.New("Username '" + newName + "' is taken already")
	}

	user.ForceSetNick(newName)
	return nil
}

// ForceSetNick forces a nickname overwrite.
func (user *User) ForceSetNick(newName string) {
	// Delete old nick reference
	DB.Delete("NickToUser", user.Nick)

	// Set new nick
	user.Nick = newName

	DB.Set("NickToUser", user.Nick, &NickToUser{
		Nick:   user.Nick,
		UserID: user.ID,
	})
}

// CleanNick only returns the nickname if it was set by user, otherwise empty string.
func (user *User) CleanNick() string {
	if user.HasNick() {
		return user.Nick
	}

	return ""
}

// Creator needs to be implemented for the PostParent interface.
func (user *User) Creator() *User {
	return user
}

// CreatorID needs to be implemented for the PostParent interface.
func (user *User) CreatorID() UserID {
	return user.ID
}

// TitleByUser returns the username.
func (user *User) TitleByUser(viewUser *User) string {
	return user.Nick
}

// SetEmail changes the user's email safely.
func (user *User) SetEmail(newEmail string) error {
	setEmailMutex.Lock()
	defer setEmailMutex.Unlock()

	if !validate.Email(newEmail) {
		return errors.New("Invalid email address")
	}

	// Delete old email reference
	DB.Delete("EmailToUser", user.Email)

	// Set new email
	user.Email = newEmail

	DB.Set("EmailToUser", user.Email, &EmailToUser{
		Email:  user.Email,
		UserID: user.ID,
	})

	return nil
}

// HasBasicInfo returns true if the user has a username, an avatar and an introduction.
func (user *User) HasBasicInfo() bool {
	return user.HasAvatar() && user.HasNick() && user.Introduction != ""
}

// TypeName returns the type name.
func (user *User) TypeName() string {
	return "User"
}

// Self returns the object itself.
func (user *User) Self() Loggable {
	return user
}
