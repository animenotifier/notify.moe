package arn

import "github.com/animenotifier/notify.moe/arn/validate"

// AnimeEpisode ...
type AnimeEpisode struct {
	Number     int               `json:"number" editable:"true"`
	Title      EpisodeTitle      `json:"title" editable:"true"`
	AiringDate AiringDate        `json:"airingDate" editable:"true"`
	Links      map[string]string `json:"links"`
}

// EpisodeTitle ...
type EpisodeTitle struct {
	Romaji   string `json:"romaji" editable:"true"`
	English  string `json:"english" editable:"true"`
	Japanese string `json:"japanese" editable:"true"`
}

// Available tells you whether the episode is available (triggered when it has a link).
func (a *AnimeEpisode) Available() bool {
	return len(a.Links) > 0
}

// AvailableOn tells you whether the episode is available on a given service.
func (a *AnimeEpisode) AvailableOn(serviceName string) bool {
	return a.Links[serviceName] != ""
}

// Merge combines the data of both episodes to one.
func (a *AnimeEpisode) Merge(b *AnimeEpisode) {
	if b == nil {
		return
	}

	a.Number = b.Number

	// Titles
	if b.Title.Romaji != "" {
		a.Title.Romaji = b.Title.Romaji
	}

	if b.Title.English != "" {
		a.Title.English = b.Title.English
	}

	if b.Title.Japanese != "" {
		a.Title.Japanese = b.Title.Japanese
	}

	// Airing date
	if validate.DateTime(b.AiringDate.Start) {
		a.AiringDate.Start = b.AiringDate.Start
	}

	if validate.DateTime(b.AiringDate.End) {
		a.AiringDate.End = b.AiringDate.End
	}

	// Links
	if a.Links == nil {
		a.Links = map[string]string{}
	}

	for name, link := range b.Links {
		a.Links[name] = link
	}
}

// NewAnimeEpisode creates an empty anime episode.
func NewAnimeEpisode() *AnimeEpisode {
	return &AnimeEpisode{
		Number:     -1,
		Title:      EpisodeTitle{},
		AiringDate: AiringDate{},
		Links:      map[string]string{},
	}
}
