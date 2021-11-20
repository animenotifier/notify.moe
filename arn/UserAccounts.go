package arn

// UserAccounts represents a user's accounts on external services.
type UserAccounts struct {
	Facebook struct {
		ID string `json:"id" private:"true"`
	} `json:"facebook"`

	Google struct {
		ID string `json:"id" private:"true"`
	} `json:"google"`

	Twitter struct {
		ID   string `json:"id" private:"true"`
		Nick string `json:"nick" private:"true"`
	} `json:"twitter"`

	Discord struct {
		Nick     string `json:"nick" editable:"true"`
		Verified bool   `json:"verified"`
	} `json:"discord"`

	AniList struct {
		Nick string `json:"nick" editable:"true"`
	} `json:"anilist"`

	AnimePlanet struct {
		Nick string `json:"nick" editable:"true"`
	} `json:"animeplanet"`

	MyAnimeList struct {
		Nick string `json:"nick" editable:"true"`
	} `json:"myanimelist"`

	Kitsu struct {
		Nick string `json:"nick" editable:"true"`
	} `json:"kitsu"`
}
