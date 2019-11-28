package arn

import (
	"fmt"
	"time"

	"github.com/aerogo/nano"
	"github.com/animenotifier/notify.moe/arn/validate"
)

// EpisodeID represents an episode ID.
type EpisodeID = ID

// Episode represents a single episode for an anime.
type Episode struct {
	AnimeID    AnimeID           `json:"animeId"`
	Number     int               `json:"number" editable:"true"`
	Title      EpisodeTitle      `json:"title" editable:"true"`
	AiringDate AiringDate        `json:"airingDate" editable:"true"`
	Links      map[string]string `json:"links"`

	hasID
}

// EpisodeTitle contains the title information for an anime episode.
type EpisodeTitle struct {
	Romaji   string `json:"romaji" editable:"true"`
	English  string `json:"english" editable:"true"`
	Japanese string `json:"japanese" editable:"true"`
}

// NewAnimeEpisode creates a new anime episode.
func NewAnimeEpisode() *Episode {
	return &Episode{
		hasID: hasID{
			ID: GenerateID("Episode"),
		},
		Number: -1,
	}
}

// Anime returns the anime the episode refers to.
func (episode *Episode) Anime() *Anime {
	anime, _ := GetAnime(episode.AnimeID)
	return anime
}

// GetID returns the episode ID.
func (episode *Episode) GetID() string {
	return episode.ID
}

// TypeName returns the type name.
func (episode *Episode) TypeName() string {
	return "Episode"
}

// Self returns the object itself.
func (episode *Episode) Self() Loggable {
	return episode
}

// Link returns the permalink to the episode.
func (episode *Episode) Link() string {
	return "/episode/" + episode.ID
}

// Available tells you whether the episode is available (triggered when it has a link).
func (episode *Episode) Available() bool {
	availableDate, err := time.Parse(time.RFC3339, episode.AiringDate.End)

	if err != nil {
		return false
	}

	return time.Now().UnixNano() > availableDate.UnixNano()
}

// Previous returns the previous episode, if available.
func (episode *Episode) Previous() *Episode {
	episodes := episode.Anime().Episodes()
	_, index := episodes.Find(episode.Number)

	if index > 0 {
		return episodes[index-1]
	}

	return nil
}

// Next returns the next episode, if available.
func (episode *Episode) Next() *Episode {
	episodes := episode.Anime().Episodes()
	_, index := episodes.Find(episode.Number)

	if index != -1 && index+1 < len(episodes) {
		return episodes[index+1]
	}

	return nil
}

// Merge combines the data of both episodes to one.
func (episode *Episode) Merge(b *Episode) {
	if b == nil {
		return
	}

	episode.Number = b.Number

	// Titles
	if b.Title.Romaji != "" {
		episode.Title.Romaji = b.Title.Romaji
	}

	if b.Title.English != "" {
		episode.Title.English = b.Title.English
	}

	if b.Title.Japanese != "" {
		episode.Title.Japanese = b.Title.Japanese
	}

	// Airing date
	if validate.DateTime(b.AiringDate.Start) {
		episode.AiringDate.Start = b.AiringDate.Start
	}

	if validate.DateTime(b.AiringDate.End) {
		episode.AiringDate.End = b.AiringDate.End
	}

	// Links
	if episode.Links == nil {
		episode.Links = map[string]string{}
	}

	for name, link := range b.Links {
		episode.Links[name] = link
	}
}

// String implements the default string serialization.
func (episode *Episode) String() string {
	return fmt.Sprintf("%s ep. %d", episode.Anime().TitleByUser(nil), episode.Number)
}

// StreamEpisodes returns a stream of all episodes.
func StreamEpisodes() <-chan *Episode {
	channel := make(chan *Episode, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("Episode") {
			channel <- obj.(*Episode)
		}

		close(channel)
	}()

	return channel
}

// GetEpisode returns the episode with the given ID.
func GetEpisode(id EpisodeID) (*Episode, error) {
	obj, err := DB.Get("Episode", id)

	if err != nil {
		return nil, err
	}

	return obj.(*Episode), nil
}
