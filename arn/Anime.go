package arn

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/aerogo/nano"
	"github.com/akyoto/color"
	"github.com/animenotifier/kitsu"
	"github.com/animenotifier/notify.moe/arn/validate"
	"github.com/animenotifier/shoboi"
	"github.com/animenotifier/twist"
)

// AnimeID represents an anime ID.
type AnimeID = ID

// Anime represents an anime.
type Anime struct {
	ID            AnimeID          `json:"id" primary:"true"`
	Type          string           `json:"type" editable:"true" datalist:"anime-types"`
	Title         MediaTitle       `json:"title" editable:"true"`
	Summary       string           `json:"summary" editable:"true" type:"textarea"`
	Status        string           `json:"status" editable:"true" datalist:"anime-status"`
	Genres        []string         `json:"genres" editable:"true"`
	StartDate     string           `json:"startDate" editable:"true"`
	EndDate       string           `json:"endDate" editable:"true"`
	EpisodeCount  int              `json:"episodeCount" editable:"true"`
	EpisodeLength int              `json:"episodeLength" editable:"true"`
	Source        string           `json:"source" editable:"true" datalist:"anime-sources"`
	Image         Image            `json:"image"`
	FirstChannel  string           `json:"firstChannel"`
	Rating        AnimeRating      `json:"rating"`
	Popularity    AnimePopularity  `json:"popularity"`
	Trailers      []*ExternalMedia `json:"trailers" editable:"true"`
	EpisodeIDs    []string         `json:"episodes"`

	// Mixins
	hasMappings
	hasPosts
	hasLikes
	hasCreator
	hasEditor
	hasDraft

	// Company IDs
	StudioIDs   []CompanyID `json:"studios" editable:"true"`
	ProducerIDs []CompanyID `json:"producers" editable:"true"`
	LicensorIDs []CompanyID `json:"licensors" editable:"true"`

	// Links to external websites
	Links []*Link `json:"links" editable:"true"`
}

// NewAnime creates a new anime.
func NewAnime() *Anime {
	anime := &Anime{}
	return anime.init()
}

// init is the constructor for Anime.
func (anime *Anime) init() *Anime {
	anime.ID = GenerateID("Anime")
	anime.Type = "tv"
	anime.Status = "upcoming"
	anime.Trailers = []*ExternalMedia{}
	anime.Mappings = []*Mapping{}
	anime.Created = DateTimeUTC()
	return anime
}

// GetAnime gets the anime with the given ID.
func GetAnime(id AnimeID) (*Anime, error) {
	obj, err := DB.Get("Anime", id)

	if err != nil {
		return nil, err
	}

	return obj.(*Anime), nil
}

// TitleByUser returns the preferred title for the given user.
func (anime *Anime) TitleByUser(user *User) string {
	return anime.Title.ByUser(user)
}

// Publish publishes the anime draft.
func (anime *Anime) Publish() error {
	// No type
	if anime.Type == "" {
		return errors.New("No type")
	}

	// No name
	if anime.Title.Canonical == "" {
		return errors.New("No canonical anime name")
	}

	// No status
	if anime.Status == "" {
		return errors.New("No status")
	}

	// No genres
	if len(anime.Genres) == 0 {
		return errors.New("No genres")
	}

	// No image
	if !anime.HasImage() {
		return errors.New("No anime image")
	}

	return publish(anime)
}

// Unpublish turns the anime into a draft.
func (anime *Anime) Unpublish() error {
	return unpublish(anime)
}

// AddStudio adds the company ID to the studio ID list if it doesn't exist already.
func (anime *Anime) AddStudio(companyID string) {
	// Is the ID valid?
	if companyID == "" {
		return
	}

	// If it already exists we don't need to add it
	for _, id := range anime.StudioIDs {
		if id == companyID {
			return
		}
	}

	anime.StudioIDs = append(anime.StudioIDs, companyID)
}

// AddProducer adds the company ID to the producer ID list if it doesn't exist already.
func (anime *Anime) AddProducer(companyID string) {
	// Is the ID valid?
	if companyID == "" {
		return
	}

	// If it already exists we don't need to add it
	for _, id := range anime.ProducerIDs {
		if id == companyID {
			return
		}
	}

	anime.ProducerIDs = append(anime.ProducerIDs, companyID)
}

// AddLicensor adds the company ID to the licensor ID list if it doesn't exist already.
func (anime *Anime) AddLicensor(companyID string) {
	// Is the ID valid?
	if companyID == "" {
		return
	}

	// If it already exists we don't need to add it
	for _, id := range anime.LicensorIDs {
		if id == companyID {
			return
		}
	}

	anime.LicensorIDs = append(anime.LicensorIDs, companyID)
}

// Studios returns the list of studios for this anime.
func (anime *Anime) Studios() []*Company {
	companies := []*Company{}

	for _, obj := range DB.GetMany("Company", anime.StudioIDs) {
		if obj == nil {
			continue
		}

		companies = append(companies, obj.(*Company))
	}

	return companies
}

// Producers returns the list of producers for this anime.
func (anime *Anime) Producers() []*Company {
	companies := []*Company{}

	for _, obj := range DB.GetMany("Company", anime.ProducerIDs) {
		if obj == nil {
			continue
		}

		companies = append(companies, obj.(*Company))
	}

	return companies
}

// Licensors returns the list of licensors for this anime.
func (anime *Anime) Licensors() []*Company {
	companies := []*Company{}

	for _, obj := range DB.GetMany("Company", anime.LicensorIDs) {
		if obj == nil {
			continue
		}

		companies = append(companies, obj.(*Company))
	}

	return companies
}

// Prequels returns the list of prequels for that anime.
func (anime *Anime) Prequels() []*Anime {
	prequels := []*Anime{}
	relations := anime.Relations()

	relations.Lock()
	defer relations.Unlock()

	for _, relation := range relations.Items {
		if relation.Type != "prequel" {
			continue
		}

		prequel := relation.Anime()

		if prequel == nil {
			color.Red("Anime %s has invalid anime relation ID %s", anime.ID, relation.AnimeID)
			continue
		}

		prequels = append(prequels, prequel)
	}

	return prequels
}

// ImageLink requires a size parameter and returns a link to the image in the given size.
func (anime *Anime) ImageLink(size string) string {
	extension := ".jpg"

	if size == "original" {
		extension = anime.Image.Extension
	}

	return fmt.Sprintf("//%s/images/anime/%s/%s%s?%v", MediaHost, size, anime.ID, extension, anime.Image.LastModified)
}

// HasImage returns whether the anime has an image or not.
func (anime *Anime) HasImage() bool {
	return anime.Image.Extension != "" && anime.Image.Width > 0
}

// AverageColor returns the average color of the image.
func (anime *Anime) AverageColor() string {
	color := anime.Image.AverageColor

	if color.Hue == 0 && color.Saturation == 0 && color.Lightness == 0 {
		return ""
	}

	return color.String()
}

// Season returns the season the anime started airing in.
func (anime *Anime) Season() string {
	if !validate.Date(anime.StartDate) && !validate.YearMonth(anime.StartDate) {
		return ""
	}

	return DateToSeason(anime.StartDateTime())
}

// Characters returns the anime characters for this anime.
func (anime *Anime) Characters() *AnimeCharacters {
	characters, _ := GetAnimeCharacters(anime.ID)
	return characters
}

// Relations ...
func (anime *Anime) Relations() *AnimeRelations {
	relations, _ := GetAnimeRelations(anime.ID)
	return relations
}

// Link returns the URI to the anime page.
func (anime *Anime) Link() string {
	return "/anime/" + anime.ID
}

// StartDateTime returns the start date as a time object.
func (anime *Anime) StartDateTime() time.Time {
	format := validate.DateFormat

	switch {
	case len(anime.StartDate) >= len(validate.DateFormat):
		// ...
	case len(anime.StartDate) >= len("2006-01"):
		format = "2006-01"
	case len(anime.StartDate) >= len("2006"):
		format = "2006"
	}

	t, _ := time.Parse(format, anime.StartDate)
	return t
}

// EndDateTime returns the end date as a time object.
func (anime *Anime) EndDateTime() time.Time {
	format := validate.DateFormat

	switch {
	case len(anime.EndDate) >= len(validate.DateFormat):
		// ...
	case len(anime.EndDate) >= len("2006-01"):
		format = "2006-01"
	case len(anime.EndDate) >= len("2006"):
		format = "2006"
	}

	t, _ := time.Parse(format, anime.EndDate)
	return t
}

// Episodes returns the anime episodes.
func (anime *Anime) Episodes() EpisodeList {
	objects := DB.GetMany("Episode", anime.EpisodeIDs)
	episodes := make([]*Episode, 0, len(anime.EpisodeIDs))

	for _, obj := range objects {
		if obj == nil {
			continue
		}

		episodes = append(episodes, obj.(*Episode))
	}

	return episodes
}

// UsersWatchingOrPlanned returns a list of users who are watching the anime right now.
func (anime *Anime) UsersWatchingOrPlanned() []*User {
	users := FilterUsers(func(user *User) bool {
		item := user.AnimeList().Find(anime.ID)

		if item == nil {
			return false
		}

		return item.Status == AnimeListStatusWatching || item.Status == AnimeListStatusPlanned
	})

	return users
}

// RefreshEpisodes will refresh the episode data.
func (anime *Anime) RefreshEpisodes() error {
	// Fetch episodes
	episodes := anime.Episodes()

	// Save number of available episodes for comparison later
	oldAvailableCount := episodes.AvailableCount()

	// Shoboi
	shoboiEpisodes, err := anime.ShoboiEpisodes()

	if err != nil {
		return err
	}

	episodes = episodes.Merge(shoboiEpisodes)

	// Cap the number of episodes
	if anime.EpisodeCount > 0 && len(episodes) > anime.EpisodeCount {
		episodes = episodes[:anime.EpisodeCount]
	}

	// Count number of available episodes
	newAvailableCount := episodes.AvailableCount()

	if anime.Status != "finished" && newAvailableCount > oldAvailableCount {
		// New episodes have been released.
		// Notify all users who are watching the anime.
		go func() {
			for _, user := range anime.UsersWatchingOrPlanned() {
				if !user.Settings().Notification.AnimeEpisodeReleases {
					continue
				}

				user.SendNotification(&PushNotification{
					Title:   anime.Title.ByUser(user),
					Message: "Episode " + strconv.Itoa(newAvailableCount) + " has been released!",
					Icon:    anime.ImageLink("medium"),
					Link:    "https://notify.moe" + anime.Link(),
					Type:    NotificationTypeAnimeEpisode,
				})
			}
		}()
	}

	// Number remaining episodes
	startNumber := 0

	for _, episode := range episodes {
		if episode.Number != -1 {
			startNumber = episode.Number
			continue
		}

		startNumber++
		episode.Number = startNumber
	}

	// Guess airing dates
	oneWeek := 7 * 24 * time.Hour
	lastAiringDate := ""
	timeDifference := oneWeek

	for _, episode := range episodes {
		if validate.DateTime(episode.AiringDate.Start) {
			if lastAiringDate != "" {
				a, _ := time.Parse(time.RFC3339, lastAiringDate)
				b, _ := time.Parse(time.RFC3339, episode.AiringDate.Start)
				timeDifference = b.Sub(a)

				// Cap time difference at one week
				if timeDifference > oneWeek {
					timeDifference = oneWeek
				}
			}

			lastAiringDate = episode.AiringDate.Start
			continue
		}

		// Add 1 week to the last known airing date
		nextAiringDate, _ := time.Parse(time.RFC3339, lastAiringDate)
		nextAiringDate = nextAiringDate.Add(timeDifference)

		// Guess start and end time
		episode.AiringDate.Start = nextAiringDate.Format(time.RFC3339)
		episode.AiringDate.End = nextAiringDate.Add(30 * time.Minute).Format(time.RFC3339)

		// Set this date as the new last known airing date
		lastAiringDate = episode.AiringDate.Start
	}

	// Save new episode ID list
	episodeIDs := make([]string, len(episodes))

	for index, episode := range episodes {
		episodeIDs[index] = episode.ID
		episode.AnimeID = anime.ID
		episode.Save()
	}

	anime.EpisodeIDs = episodeIDs
	anime.Save()
	return nil
}

// ShoboiEpisodes returns a slice of episode info from cal.syoboi.jp.
func (anime *Anime) ShoboiEpisodes() (EpisodeList, error) {
	shoboiID := anime.GetMapping("shoboi/anime")

	if shoboiID == "" {
		return nil, errors.New("Missing shoboi/anime mapping")
	}

	shoboiAnime, err := shoboi.GetAnime(shoboiID)

	if err != nil {
		return nil, err
	}

	arnEpisodes := []*Episode{}
	shoboiEpisodes := shoboiAnime.Episodes()

	for _, shoboiEpisode := range shoboiEpisodes {
		episode := NewAnimeEpisode()
		episode.Number = shoboiEpisode.Number
		episode.Title.Japanese = shoboiEpisode.TitleJapanese

		// Try to get airing date
		airingDate := shoboiEpisode.AiringDate

		if airingDate != nil {
			episode.AiringDate.Start = airingDate.Start
			episode.AiringDate.End = airingDate.End
		} else {
			episode.AiringDate.Start = ""
			episode.AiringDate.End = ""
		}

		arnEpisodes = append(arnEpisodes, episode)
	}

	return arnEpisodes, nil
}

// TwistEpisodes returns a slice of episode info from twist.moe.
func (anime *Anime) TwistEpisodes() (EpisodeList, error) {
	idList, err := GetIDList("animetwist index")

	if err != nil {
		return nil, err
	}

	// Does the index contain the ID?
	kitsuID := anime.GetMapping("kitsu/anime")
	found := false

	for _, id := range idList {
		if id == kitsuID {
			found = true
			break
		}
	}

	// If the ID is not the index we don't need to query the feed
	if !found {
		return nil, errors.New("Not available in twist.moe anime index")
	}

	// Get twist.moe feed
	feed, err := twist.GetFeedByKitsuID(kitsuID)

	if err != nil {
		return nil, err
	}

	episodes := feed.Episodes

	// Sort by episode number
	sort.Slice(episodes, func(a, b int) bool {
		return episodes[a].Number < episodes[b].Number
	})

	arnEpisodes := []*Episode{}

	for _, episode := range episodes {
		arnEpisode := NewAnimeEpisode()
		arnEpisode.Number = episode.Number
		arnEpisode.Links = map[string]string{
			"twist.moe": strings.Replace(episode.Link, "https://test.twist.moe/", "https://twist.moe/", 1),
		}

		arnEpisodes = append(arnEpisodes, arnEpisode)
	}

	return arnEpisodes, nil
}

// UpcomingEpisodes ...
func (anime *Anime) UpcomingEpisodes() []*UpcomingEpisode {
	var upcomingEpisodes []*UpcomingEpisode
	now := time.Now().UTC().Format(time.RFC3339)

	for _, episode := range anime.Episodes() {
		if episode.AiringDate.Start > now && validate.DateTime(episode.AiringDate.Start) {
			upcomingEpisodes = append(upcomingEpisodes, &UpcomingEpisode{
				Anime:   anime,
				Episode: episode,
			})
		}
	}

	return upcomingEpisodes
}

// UpcomingEpisode ...
func (anime *Anime) UpcomingEpisode() *UpcomingEpisode {
	now := time.Now().UTC().Format(time.RFC3339)

	for _, episode := range anime.Episodes() {
		if episode.AiringDate.Start > now && validate.DateTime(episode.AiringDate.Start) {
			return &UpcomingEpisode{
				Anime:   anime,
				Episode: episode,
			}
		}
	}

	return nil
}

// EpisodeCountString formats the episode count and displays
// a question mark when the number of episodes is unknown.
func (anime *Anime) EpisodeCountString() string {
	if anime.EpisodeCount == 0 {
		return "?"
	}

	return strconv.Itoa(anime.EpisodeCount)
}

// ImportKitsuMapping imports the given Kitsu mapping.
func (anime *Anime) ImportKitsuMapping(mapping *kitsu.Mapping) {
	switch mapping.Attributes.ExternalSite {
	case "myanimelist/anime":
		anime.SetMapping("myanimelist/anime", mapping.Attributes.ExternalID)
	case "anidb":
		anime.SetMapping("anidb/anime", mapping.Attributes.ExternalID)
	case "trakt":
		anime.SetMapping("trakt/anime", mapping.Attributes.ExternalID)
	// case "hulu":
	// 	anime.SetMapping("hulu/anime", mapping.Attributes.ExternalID)
	case "anilist":
		externalID := mapping.Attributes.ExternalID
		externalID = strings.TrimPrefix(externalID, "anime/")

		anime.SetMapping("anilist/anime", externalID)
	case "thetvdb", "thetvdb/series":
		externalID := mapping.Attributes.ExternalID
		slashPos := strings.Index(externalID, "/")

		if slashPos != -1 {
			externalID = externalID[:slashPos]
		}

		anime.SetMapping("thetvdb/anime", externalID)
	case "thetvdb/season":
		// Ignore
	default:
		color.Yellow("Unknown mapping: %s %s", mapping.Attributes.ExternalSite, mapping.Attributes.ExternalID)
	}
}

// TypeHumanReadable ...
func (anime *Anime) TypeHumanReadable() string {
	switch anime.Type {
	case "tv":
		return "TV"
	case "movie":
		return "Movie"
	case "ova":
		return "OVA"
	case "ona":
		return "ONA"
	case "special":
		return "Special"
	case "music":
		return "Music"
	default:
		return anime.Type
	}
}

// StatusHumanReadable ...
func (anime *Anime) StatusHumanReadable() string {
	switch anime.Status {
	case "finished":
		return "Finished"
	case "current":
		return "Airing"
	case "upcoming":
		return "Upcoming"
	case "tba":
		return "To be announced"
	default:
		return anime.Status
	}
}

// CalculatedStatus returns the status of the anime inferred by the start and end date.
func (anime *Anime) CalculatedStatus() string {
	// If we are past the end date, the anime is finished.
	if validate.Date(anime.EndDate) {
		end := anime.EndDateTime()

		if time.Since(end) > 0 {
			return "finished"
		}
	}

	// If we have a start date and we didn't reach the end date, it's either current or upcoming.
	if validate.Date(anime.StartDate) {
		start := anime.StartDateTime()

		if time.Since(start) > 0 {
			return "current"
		}

		return "upcoming"
	}

	// If we have no date information it's to be announced.
	return "tba"
}

// EpisodeByNumber returns the episode with the given number.
func (anime *Anime) EpisodeByNumber(number int) *Episode {
	for _, episode := range anime.Episodes() {
		if number == episode.Number {
			return episode
		}
	}

	return nil
}

// String implements the default string serialization.
func (anime *Anime) String() string {
	return anime.Title.Canonical
}

// GetID returns the ID.
func (anime *Anime) GetID() string {
	return anime.ID
}

// TypeName returns the type name.
func (anime *Anime) TypeName() string {
	return "Anime"
}

// Self returns the object itself.
func (anime *Anime) Self() Loggable {
	return anime
}

// StreamAnime returns a stream of all anime.
func StreamAnime() <-chan *Anime {
	channel := make(chan *Anime, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("Anime") {
			channel <- obj.(*Anime)
		}

		close(channel)
	}()

	return channel
}

// AllAnime returns a slice of all anime.
func AllAnime() []*Anime {
	all := make([]*Anime, 0, DB.Collection("Anime").Count())

	stream := StreamAnime()

	for obj := range stream {
		all = append(all, obj)
	}

	return all
}

// FilterAnime filters all anime by a custom function.
func FilterAnime(filter func(*Anime) bool) []*Anime {
	var filtered []*Anime

	channel := DB.All("Anime")

	for obj := range channel {
		realObject := obj.(*Anime)

		if filter(realObject) {
			filtered = append(filtered, realObject)
		}
	}

	return filtered
}
