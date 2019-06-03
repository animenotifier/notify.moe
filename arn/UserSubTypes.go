package arn

// UserBrowser ...
type UserBrowser struct {
	Name     string `json:"name" private:"true"`
	Version  string `json:"version" private:"true"`
	IsMobile bool   `json:"isMobile" private:"true"`
}

// UserOS ...
type UserOS struct {
	Name    string `json:"name" private:"true"`
	Version string `json:"version" private:"true"`
}

// UserListProviders ...
type UserListProviders struct {
	AniList     ListProviderConfig `json:"AniList"`
	AnimePlanet ListProviderConfig `json:"AnimePlanet"`
	HummingBird ListProviderConfig `json:"HummingBird"`
	MyAnimeList ListProviderConfig `json:"MyAnimeList"`
}

// ListProviderConfig ...
type ListProviderConfig struct {
	UserName string `json:"userName"`
}

// PushEndpoint ...
type PushEndpoint struct {
	Registered string `json:"registered"`
	Keys       struct {
		P256DH string `json:"p256dh" private:"true"`
		Auth   string `json:"auth" private:"true"`
	} `json:"keys"`
}

// CSSPosition ...
type CSSPosition struct {
	X string `json:"x"`
	Y string `json:"y"`
}
