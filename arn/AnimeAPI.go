package arn

import (
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
	"github.com/akyoto/color"
)

// Force interface implementations
var (
	_ fmt.Stringer           = (*Anime)(nil)
	_ Likeable               = (*Anime)(nil)
	_ PostParent             = (*Anime)(nil)
	_ api.Deletable          = (*Anime)(nil)
	_ api.Editable           = (*Anime)(nil)
	_ api.CustomEditable     = (*Anime)(nil)
	_ api.ArrayEventListener = (*Anime)(nil)
)

// Actions
func init() {
	API.RegisterActions("Anime", []*api.Action{
		// Like anime
		LikeAction(),

		// Unlike anime
		UnlikeAction(),
	})
}

// Edit creates an edit log entry.
func (anime *Anime) Edit(ctx aero.Context, key string, value reflect.Value, newValue reflect.Value) (consumed bool, err error) {
	user := GetUserFromContext(ctx)

	if key == "Status" {
		oldStatus := value.String()
		newStatus := newValue.String()

		// Notify people who want to know about finished series
		if oldStatus == "current" && newStatus == "finished" {
			go func() {
				for _, user := range anime.UsersWatchingOrPlanned() {
					if !user.Settings().Notification.AnimeFinished {
						continue
					}

					user.SendNotification(&PushNotification{
						Title:   anime.Title.ByUser(user),
						Message: anime.Title.ByUser(user) + " has finished airing!",
						Icon:    anime.ImageLink("medium"),
						Link:    "https://notify.moe" + anime.Link(),
						Type:    NotificationTypeAnimeFinished,
					})
				}
			}()
		}
	}

	// Write log entry
	logEntry := NewEditLogEntry(user.ID, "edit", "Anime", anime.ID, key, fmt.Sprint(value.Interface()), fmt.Sprint(newValue.Interface()))
	logEntry.Save()

	return false, nil
}

// OnAppend saves a log entry.
func (anime *Anime) OnAppend(ctx aero.Context, key string, index int, obj interface{}) {
	onAppend(anime, ctx, key, index, obj)
}

// OnRemove saves a log entry.
func (anime *Anime) OnRemove(ctx aero.Context, key string, index int, obj interface{}) {
	onRemove(anime, ctx, key, index, obj)
}

// Authorize returns an error if the given API POST request is not authorized.
func (anime *Anime) Authorize(ctx aero.Context, action string) error {
	user := GetUserFromContext(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return errors.New("Not logged in or not authorized to edit this anime")
	}

	return nil
}

// DeleteInContext deletes the anime in the given context.
func (anime *Anime) DeleteInContext(ctx aero.Context) error {
	user := GetUserFromContext(ctx)

	// Write log entry
	logEntry := NewEditLogEntry(user.ID, "delete", "Anime", anime.ID, "", fmt.Sprint(anime), "")
	logEntry.Save()

	return anime.Delete()
}

// Delete deletes the anime from the database.
func (anime *Anime) Delete() error {
	// Delete anime characters
	DB.Delete("AnimeCharacters", anime.ID)

	// Delete anime relations
	DB.Delete("AnimeRelations", anime.ID)

	// Delete anime episodes
	DB.Delete("AnimeEpisodes", anime.ID)

	// Delete anime list items
	for animeList := range StreamAnimeLists() {
		removed := animeList.Remove(anime.ID)

		if removed {
			animeList.Save()
		}
	}

	// Delete anime ID from existing relations
	for relations := range StreamAnimeRelations() {
		removed := relations.Remove(anime.ID)

		if removed {
			relations.Save()
		}
	}

	// Delete anime ID from quotes
	for quote := range StreamQuotes() {
		if quote.AnimeID == anime.ID {
			quote.AnimeID = ""
			quote.Save()
		}
	}

	// Remove posts
	for _, post := range anime.Posts() {
		err := post.Delete()

		if err != nil {
			return err
		}
	}

	// Delete soundtrack tags
	for track := range StreamSoundTracks() {
		newTags := []string{}

		for _, tag := range track.Tags {
			if strings.HasPrefix(tag, "anime:") {
				parts := strings.Split(tag, ":")
				id := parts[1]

				if id == anime.ID {
					continue
				}
			}

			newTags = append(newTags, tag)
		}

		if len(track.Tags) != len(newTags) {
			track.Tags = newTags
			track.Save()
		}
	}

	// Delete images on file system
	if anime.HasImage() {
		err := os.Remove(path.Join(Root, "images/anime/original/", anime.ID+anime.Image.Extension))

		if err != nil {
			// Don't return the error.
			// It's too late to stop the process at this point.
			// Instead, log the error.
			color.Red(err.Error())
		}

		os.Remove(path.Join(Root, "images/anime/large/", anime.ID+".jpg"))
		os.Remove(path.Join(Root, "images/anime/large/", anime.ID+"@2.jpg"))
		os.Remove(path.Join(Root, "images/anime/large/", anime.ID+".webp"))
		os.Remove(path.Join(Root, "images/anime/large/", anime.ID+"@2.webp"))

		os.Remove(path.Join(Root, "images/anime/medium/", anime.ID+".jpg"))
		os.Remove(path.Join(Root, "images/anime/medium/", anime.ID+"@2.jpg"))
		os.Remove(path.Join(Root, "images/anime/medium/", anime.ID+".webp"))
		os.Remove(path.Join(Root, "images/anime/medium/", anime.ID+"@2.webp"))

		os.Remove(path.Join(Root, "images/anime/small/", anime.ID+".jpg"))
		os.Remove(path.Join(Root, "images/anime/small/", anime.ID+"@2.jpg"))
		os.Remove(path.Join(Root, "images/anime/small/", anime.ID+".webp"))
		os.Remove(path.Join(Root, "images/anime/small/", anime.ID+"@2.webp"))
	}

	// Delete the actual anime
	DB.Delete("Anime", anime.ID)

	return nil
}

// Save saves the anime in the database.
func (anime *Anime) Save() {
	DB.Set("Anime", anime.ID, anime)
}
