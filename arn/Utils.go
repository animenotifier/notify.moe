package arn

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aerogo/aero"
	"github.com/aerogo/mirror"
	"github.com/akyoto/color"
	"github.com/animenotifier/kitsu"
	"github.com/animenotifier/mal"
	jsoniter "github.com/json-iterator/go"
	shortid "github.com/ventu-io/go-shortid"
)

var (
	// MediaHost is the host we use to link image files.
	MediaHost = "media.notify.moe"

	// Regular expressions
	stripTagsRegex = regexp.MustCompile(`<[^>]*>`)
	sourceRegex    = regexp.MustCompile(`\(Source: (.*?)\)`)
	writtenByRegex = regexp.MustCompile(`\[Written by (.*?)\]`)
)

// GenerateID generates a unique ID for a given collection.
func GenerateID(collection string) string {
	id, _ := shortid.Generate()

	// Retry until we find an unused ID
	retry := 0

	for {
		_, err := DB.Get(collection, id)

		if err != nil && strings.Contains(err.Error(), "not found") {
			return id
		}

		retry++

		if retry > 10 {
			panic(errors.New("Can't generate unique ID"))
		}

		id, _ = shortid.Generate()
	}
}

// GetUserFromContext returns the logged in user for the given context.
func GetUserFromContext(ctx aero.Context) *User {
	if !ctx.HasSession() {
		return nil
	}

	userID := ctx.Session().GetString("userId")

	if userID == "" {
		return nil
	}

	user, err := GetUser(userID)

	if err != nil {
		return nil
	}

	return user
}

// GetObjectTitle ...
func GetObjectTitle(typeName string, id string) string {
	obj, err := DB.Get(typeName, id)

	if err != nil {
		return fmt.Sprintf("<not found: %s>", id)
	}

	return fmt.Sprint(obj)
}

// GetObjectLink ...
func GetObjectLink(typeName string, id string) string {
	obj, err := DB.Get(typeName, id)

	if err != nil {
		return fmt.Sprintf("<not found: %s>", id)
	}

	linkable, ok := obj.(Linkable)

	if ok {
		return linkable.Link()
	}

	return "/" + strings.ToLower(typeName) + "/" + id
}

// FilterIDTags returns all IDs of the given type in the tag list.
func FilterIDTags(tags []string, idType string) []string {
	var idList []string
	prefix := idType + ":"

	for _, tag := range tags {
		if strings.HasPrefix(tag, prefix) {
			id := strings.TrimPrefix(tag, prefix)
			idList = append(idList, id)
		}
	}

	return idList
}

// AgeInYears returns the person's age in years.
func AgeInYears(birthDayString string) int {
	birthDay, err := time.Parse("2006-01-02", birthDayString)

	if err != nil {
		return 0
	}

	now := time.Now()
	years := now.Year() - birthDay.Year()

	if now.YearDay() < birthDay.YearDay() {
		years--
	}

	return years
}

// JSON turns the object into a JSON string.
func JSON(obj interface{}) string {
	data, err := jsoniter.Marshal(obj)

	if err == nil {
		return string(data)
	}

	return err.Error()
}

// SetObjectProperties updates the object with the given map[string]interface{}
func SetObjectProperties(rootObj interface{}, updates map[string]interface{}) error {
	for key, value := range updates {
		field, _, v, err := mirror.GetField(rootObj, key)

		if err != nil {
			return err
		}

		// Is somebody attempting to edit fields that aren't editable?
		if field.Tag.Get("editable") != "true" {
			return errors.New("Field " + key + " is not editable")
		}

		newValue := reflect.ValueOf(value)

		// Implement special data type cases here
		if v.Kind() == reflect.Int {
			x := int64(newValue.Float())

			if !v.OverflowInt(x) {
				v.SetInt(x)
			}
		} else {
			v.Set(newValue)
		}
	}

	return nil
}

// GetGenreIDByName ...
func GetGenreIDByName(genre string) string {
	genre = strings.Replace(genre, "-", "", -1)
	genre = strings.Replace(genre, " ", "", -1)
	genre = strings.ToLower(genre)
	return genre
}

// FixAnimeDescription ...
func FixAnimeDescription(description string) string {
	description = stripTagsRegex.ReplaceAllString(description, "")
	description = sourceRegex.ReplaceAllString(description, "")
	description = writtenByRegex.ReplaceAllString(description, "")
	return strings.TrimSpace(description)
}

// FixGender ...
func FixGender(gender string) string {
	if gender != "male" && gender != "female" {
		return ""
	}

	return gender
}

// DateToSeason returns the season of the year for the given date.
func DateToSeason(date time.Time) string {
	month := date.Month()

	if month >= 4 && month <= 6 {
		return "spring"
	}

	if month >= 7 && month <= 9 {
		return "summer"
	}

	if month >= 10 && month <= 12 {
		return "autumn"
	}

	if month >= 1 && month < 4 {
		return "winter"
	}

	return ""
}

// BroadcastEvent sends the given event to the event streams of all users.
func BroadcastEvent(event *aero.Event) {
	for user := range StreamUsers() {
		user.BroadcastEvent(event)
	}
}

// AnimeRatingStars displays the rating in Unicode stars.
func AnimeRatingStars(rating float64) string {
	stars := int(rating/20 + 0.5)
	return strings.Repeat("★", stars) + strings.Repeat("☆", 5-stars)
}

// EpisodesToString shows a question mark if the episode count is zero.
func EpisodesToString(episodes int) string {
	if episodes == 0 {
		return "?"
	}

	return fmt.Sprint(episodes)
}

// EpisodeCountMax is used for the max value of number input on episodes.
func EpisodeCountMax(episodes int) string {
	if episodes == 0 {
		return ""
	}

	return strconv.Itoa(episodes)
}

// DateTimeUTC returns the current UTC time in RFC3339 format.
func DateTimeUTC() string {
	return time.Now().UTC().Format(time.RFC3339)
}

// OverallRatingName returns Overall in general, but Hype when episodes watched is zero.
func OverallRatingName(episodes int) string {
	if episodes == 0 {
		return "Hype"
	}

	return "Overall"
}

// IsIPv6 tells you whether the given address is IPv6 encoded.
func IsIPv6(ip string) bool {
	for i := 0; i < len(ip); i++ {
		if ip[i] == ':' {
			return true
		}
	}

	return false
}

// MyAnimeListStatusToARNStatus ...
func MyAnimeListStatusToARNStatus(status int) string {
	switch status {
	case mal.AnimeListStatusCompleted:
		return AnimeListStatusCompleted
	case mal.AnimeListStatusWatching:
		return AnimeListStatusWatching
	case mal.AnimeListStatusPlanned:
		return AnimeListStatusPlanned
	case mal.AnimeListStatusHold:
		return AnimeListStatusHold
	case mal.AnimeListStatusDropped:
		return AnimeListStatusDropped
	default:
		return ""
	}
}

// KitsuStatusToARNStatus ...
func KitsuStatusToARNStatus(status string) string {
	switch status {
	case kitsu.AnimeListStatusCompleted:
		return AnimeListStatusCompleted
	case kitsu.AnimeListStatusWatching:
		return AnimeListStatusWatching
	case kitsu.AnimeListStatusPlanned:
		return AnimeListStatusPlanned
	case kitsu.AnimeListStatusHold:
		return AnimeListStatusHold
	case kitsu.AnimeListStatusDropped:
		return AnimeListStatusDropped
	default:
		return ""
	}
}

// ListItemStatusName ...
func ListItemStatusName(status string) string {
	switch status {
	case AnimeListStatusWatching:
		return "Watching"
	case AnimeListStatusCompleted:
		return "Completed"
	case AnimeListStatusPlanned:
		return "Planned"
	case AnimeListStatusHold:
		return "On hold"
	case AnimeListStatusDropped:
		return "Dropped"
	default:
		return ""
	}
}

// IsTest returns true if the program is currently running in the "go test" tool.
func IsTest() bool {
	return flag.Lookup("test.v") != nil
}

// PanicOnError will panic if the error is not nil.
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// deleteImages deletes images in the given folder.
func deleteImages(folderName string, id string, originalExtension string) {
	if originalExtension == "" {
		return
	}

	err := os.Remove(path.Join(Root, "images", folderName, "original", id+originalExtension))

	if err != nil {
		// Don't return the error.
		// It's too late to stop the process at this point.
		// Instead, log the error.
		color.Red(err.Error())
	}

	os.Remove(path.Join(Root, "images", folderName, "large", id+".jpg"))
	os.Remove(path.Join(Root, "images", folderName, "large", id+"@2.jpg"))
	os.Remove(path.Join(Root, "images", folderName, "large", id+".webp"))
	os.Remove(path.Join(Root, "images", folderName, "large", id+"@2.webp"))
	os.Remove(path.Join(Root, "images", folderName, "medium", id+".jpg"))
	os.Remove(path.Join(Root, "images", folderName, "medium", id+"@2.jpg"))
	os.Remove(path.Join(Root, "images", folderName, "medium", id+".webp"))
	os.Remove(path.Join(Root, "images", folderName, "medium", id+"@2.webp"))
	os.Remove(path.Join(Root, "images", folderName, "small", id+".jpg"))
	os.Remove(path.Join(Root, "images", folderName, "small", id+"@2.jpg"))
	os.Remove(path.Join(Root, "images", folderName, "small", id+".webp"))
	os.Remove(path.Join(Root, "images", folderName, "small", id+"@2.webp"))
}
