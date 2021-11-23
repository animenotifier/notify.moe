package arn

import "fmt"

const (
	// SortByAiringDate sorts your watching list by airing date.
	SortByAiringDate = "airing date"

	// SortByTitle sorts your watching list alphabetically.
	SortByTitle = "title"

	// SortByRating sorts your watching list by rating.
	SortByRating = "rating"
)

const (
	// TitleLanguageCanonical ...
	TitleLanguageCanonical = "canonical"

	// TitleLanguageRomaji ...
	TitleLanguageRomaji = "romaji"

	// TitleLanguageEnglish ...
	TitleLanguageEnglish = "english"

	// TitleLanguageJapanese ...
	TitleLanguageJapanese = "japanese"
)

// Settings represents user settings.
type Settings struct {
	UserID        UserID               `json:"userId" primary:"true"`
	SortBy        string               `json:"sortBy" editable:"true"`
	TitleLanguage string               `json:"titleLanguage" editable:"true"`
	Providers     ServiceProviders     `json:"providers"`
	Format        FormatSettings       `json:"format"`
	Notification  NotificationSettings `json:"notification"`
	Editor        EditorSettings       `json:"editor"`
	Privacy       PrivacySettings      `json:"privacy"`
	Activity      ActivitySettings     `json:"activity"`
	Calendar      CalendarSettings     `json:"calendar" editable:"true"`
	Theme         string               `json:"theme" editable:"true"`
}

// ActivitySettings ...
type ActivitySettings struct {
	ShowFollowedOnly bool `json:"showFollowedOnly" editable:"true"`
}

// PrivacySettings ...
type PrivacySettings struct {
	ShowAge      bool `json:"showAge" editable:"true"`
	ShowGender   bool `json:"showGender" editable:"true"`
	ShowLocation bool `json:"showLocation" editable:"true"`
}

// NotificationSettings ...
type NotificationSettings struct {
	Email                string `json:"email" private:"true"`
	NewFollowers         bool   `json:"newFollowers" editable:"true"`
	AnimeEpisodeReleases bool   `json:"animeEpisodeReleases" editable:"true"`
	AnimeFinished        bool   `json:"animeFinished" editable:"true"`
	ForumLikes           bool   `json:"forumLikes" editable:"true"`
	GroupPostLikes       bool   `json:"groupPostLikes" editable:"true"`
	QuoteLikes           bool   `json:"quoteLikes" editable:"true"`
	SoundTrackLikes      bool   `json:"soundTrackLikes" editable:"true"`
}

// EditorSettings ...
type EditorSettings struct {
	Filter EditorFilterSettings `json:"filter"`
}

// EditorFilterSettings ...
type EditorFilterSettings struct {
	Year   string `json:"year" editable:"true"`
	Season string `json:"season" editable:"true"`
	Status string `json:"status" editable:"true"`
	Type   string `json:"type" editable:"true"`
}

// Suffix returns the URL suffix.
func (filter *EditorFilterSettings) Suffix() string {
	year := filter.Year
	status := filter.Status
	season := filter.Season
	typ := filter.Type

	if year == "" {
		year = "any"
	}

	if season == "" {
		season = "any"
	}

	if status == "" {
		status = "any"
	}

	if typ == "" {
		typ = "any"
	}

	return fmt.Sprintf("/%s/%s/%s/%s", year, season, status, typ)
}

// FormatSettings ...
type FormatSettings struct {
	RatingsPrecision int `json:"ratingsPrecision" editable:"true"`
}

// ServiceProviders ...
type ServiceProviders struct {
	Anime string `json:"anime"`
}

// CalendarSettings ...
type CalendarSettings struct {
	ShowAddedAnimeOnly bool `json:"showAddedAnimeOnly" editable:"true"`
}

// NewSettings creates the default settings for a new user.
func NewSettings(user *User) *Settings {
	return &Settings{
		UserID:        user.ID,
		SortBy:        SortByRating,
		TitleLanguage: TitleLanguageCanonical,
		Providers: ServiceProviders{
			Anime: "",
		},
		Format: FormatSettings{
			RatingsPrecision: 1,
		},
		Privacy: PrivacySettings{
			ShowLocation: true,
		},
		Calendar: CalendarSettings{
			ShowAddedAnimeOnly: false,
		},
		Notification: DefaultNotificationSettings(),
		Theme:        "light",
	}
}

// DefaultNotificationSettings returns the default notification settings.
func DefaultNotificationSettings() NotificationSettings {
	return NotificationSettings{
		Email:                "",
		NewFollowers:         true,
		AnimeEpisodeReleases: true,
		AnimeFinished:        false,
		ForumLikes:           true,
		GroupPostLikes:       true,
		QuoteLikes:           true,
		SoundTrackLikes:      true,
	}
}

// GetSettings ...
func GetSettings(userID UserID) (*Settings, error) {
	obj, err := DB.Get("Settings", userID)

	if err != nil {
		return nil, err
	}

	return obj.(*Settings), nil
}

// GetID returns the ID.
func (settings *Settings) GetID() string {
	return settings.UserID
}

// User returns the user object for the settings.
func (settings *Settings) User() *User {
	user, _ := GetUser(settings.UserID)
	return user
}
