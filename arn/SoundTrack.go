package arn

import (
	"errors"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"

	"github.com/aerogo/nano"
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn/autocorrect"
)

// SoundTrack is a soundtrack used in one or multiple anime.
type SoundTrack struct {
	Title  SoundTrackTitle  `json:"title" editable:"true"`
	Media  []*ExternalMedia `json:"media" editable:"true"`
	Links  []*Link          `json:"links" editable:"true"`
	Lyrics SoundTrackLyrics `json:"lyrics" editable:"true"`
	Tags   []string         `json:"tags" editable:"true" tooltip:"<ul><li><strong>anime:ID</strong> to connect it with anime (e.g. anime:yF1RhKiiR)</li><li><strong>opening</strong> for openings</li><li><strong>ending</strong> for endings</li><li><strong>op:NUMBER</strong> or <strong>ed:NUMBER</strong> if it has more than one OP/ED (e.g. op:2 or ed:3)</li><li><strong>cover</strong> for covers</li><li><strong>remix</strong> for remixes</li><li><strong>male</strong> or <strong>female</strong></li><li><strong title='Has lyrics'>vocal</strong>, <strong title='Has orchestral instruments, mostly no lyrics'>orchestral</strong> or <strong title='Has a mix of different instruments, mostly no lyrics'>instrumental</strong></li></ul>"`
	File   string           `json:"file"`

	hasID
	hasPosts
	hasCreator
	hasEditor
	hasLikes
	hasDraft
}

// Link returns the permalink for the track.
func (track *SoundTrack) Link() string {
	return "/soundtrack/" + track.ID
}

// TitleByUser returns the preferred title for the given user.
func (track *SoundTrack) TitleByUser(user *User) string {
	return track.Title.ByUser(user)
}

// MediaByService returns a slice of all media by the given service.
func (track *SoundTrack) MediaByService(service string) []*ExternalMedia {
	filtered := []*ExternalMedia{}

	for _, media := range track.Media {
		if media.Service == service {
			filtered = append(filtered, media)
		}
	}

	return filtered
}

// HasMediaByService returns true if the track has media by the given service.
func (track *SoundTrack) HasMediaByService(service string) bool {
	for _, media := range track.Media {
		if media.Service == service {
			return true
		}
	}

	return false
}

// HasTag returns true if it contains the given tag.
func (track *SoundTrack) HasTag(search string) bool {
	for _, tag := range track.Tags {
		if tag == search {
			return true
		}
	}

	return false
}

// HasLyrics returns true if the track has lyrics in any language.
func (track *SoundTrack) HasLyrics() bool {
	return track.Lyrics.Native != "" || track.Lyrics.Romaji != ""
}

// Anime fetches all tagged anime of the sound track.
func (track *SoundTrack) Anime() []*Anime {
	var animeList []*Anime

	for _, tag := range track.Tags {
		if strings.HasPrefix(tag, "anime:") {
			animeID := strings.TrimPrefix(tag, "anime:")
			anime, err := GetAnime(animeID)

			if err != nil {
				if !track.IsDraft {
					color.Red("Error fetching anime: %v", err)
				}

				continue
			}

			animeList = append(animeList, anime)
		}
	}

	return animeList
}

// MainAnime ...
func (track *SoundTrack) MainAnime() *Anime {
	allAnime := track.Anime()

	if len(allAnime) == 0 {
		return nil
	}

	return allAnime[0]
}

// TypeName returns the type name.
func (track *SoundTrack) TypeName() string {
	return "SoundTrack"
}

// Self returns the object itself.
func (track *SoundTrack) Self() Loggable {
	return track
}

// EditedByUser returns the user who edited this track last.
func (track *SoundTrack) EditedByUser() *User {
	user, _ := GetUser(track.EditedBy)
	return user
}

// OnLike is called when the soundtrack receives a like.
func (track *SoundTrack) OnLike(likedBy *User) {
	if likedBy.ID == track.CreatedBy {
		return
	}

	if !track.Creator().Settings().Notification.SoundTrackLikes {
		return
	}

	go func() {
		track.Creator().SendNotification(&PushNotification{
			Title:   likedBy.Nick + " liked your soundtrack " + track.Title.ByUser(track.Creator()),
			Message: likedBy.Nick + " liked your soundtrack " + track.Title.ByUser(track.Creator()) + ".",
			Icon:    "https:" + likedBy.AvatarLink("large"),
			Link:    "https://notify.moe" + likedBy.Link(),
			Type:    NotificationTypeLike,
		})
	}()
}

// Publish ...
func (track *SoundTrack) Publish() error {
	// No media added
	if len(track.Media) == 0 {
		return errors.New("No media specified (at least 1 media source is required)")
	}

	animeFound := false

	for _, tag := range track.Tags {
		tag = autocorrect.Tag(tag)

		if strings.HasPrefix(tag, "anime:") {
			animeID := strings.TrimPrefix(tag, "anime:")
			_, err := GetAnime(animeID)

			if err != nil {
				return errors.New("Invalid anime ID")
			}

			animeFound = true
		}
	}

	// No anime found
	if !animeFound {
		return errors.New("Need to specify at least one anime")
	}

	// No tags
	if len(track.Tags) < 1 {
		return errors.New("Need to specify at least one tag")
	}

	// Publish
	err := publish(track)

	if err != nil {
		return err
	}

	// Start download in the background
	go func() {
		err := track.Download()

		if err == nil {
			track.Save()
		}
	}()

	return nil
}

// Unpublish ...
func (track *SoundTrack) Unpublish() error {
	draftIndex, err := GetDraftIndex(track.CreatedBy)

	if err != nil {
		return err
	}

	if draftIndex.SoundTrackID != "" {
		return errors.New("You still have an unfinished draft")
	}

	track.IsDraft = true
	draftIndex.SoundTrackID = track.ID
	draftIndex.Save()
	return nil
}

// Download downloads the track.
func (track *SoundTrack) Download() error {
	if track.IsDraft {
		return errors.New("Track is a draft")
	}

	youtubeVideos := track.MediaByService("Youtube")

	if len(youtubeVideos) == 0 {
		return errors.New("No Youtube ID")
	}

	youtubeID := youtubeVideos[0].ServiceID

	// Check for existing file
	if track.File != "" {
		stat, err := os.Stat(path.Join(Root, "audio", track.File))

		if err == nil && !stat.IsDir() && stat.Size() > 0 {
			return errors.New("Already downloaded")
		}
	}

	audioDirectory := path.Join(Root, "audio")
	baseName := track.ID + "|" + youtubeID

	// Check if it exists on the file system
	fullPath := FindFileWithExtension(baseName, audioDirectory, []string{
		".opus",
		".webm",
		".ogg",
		".m4a",
		".mp3",
		".flac",
		".wav",
	})

	// In case we added the file but didn't register it in database
	if fullPath != "" {
		extension := path.Ext(fullPath)
		track.File = baseName + extension
		return nil
	}

	filePath := path.Join(audioDirectory, baseName)

	// Use full URL to avoid problems with Youtube IDs that start with a hyphen
	url := "https://youtube.com/watch?v=" + youtubeID

	// Download
	cmd := exec.Command(
		"youtube-dl",
		"--no-check-certificate",
		"--extract-audio",
		"--audio-quality", "0",
		"--output", filePath+".%(ext)s",
		url,
	)

	err := cmd.Start()

	if err != nil {
		return err
	}

	err = cmd.Wait()

	if err != nil {
		return err
	}

	// Find downloaded file
	fullPath = FindFileWithExtension(baseName, audioDirectory, []string{
		".opus",
		".webm",
		".ogg",
		".m4a",
		".mp3",
		".flac",
		".wav",
	})

	extension := path.Ext(fullPath)
	track.File = baseName + extension
	return nil
}

// String implements the default string serialization.
func (track *SoundTrack) String() string {
	return track.Title.ByUser(nil)
}

// SortSoundTracksLatestFirst ...
func SortSoundTracksLatestFirst(tracks []*SoundTrack) {
	sort.Slice(tracks, func(i, j int) bool {
		return tracks[i].Created > tracks[j].Created
	})
}

// SortSoundTracksPopularFirst ...
func SortSoundTracksPopularFirst(tracks []*SoundTrack) {
	sort.Slice(tracks, func(i, j int) bool {
		aLikes := len(tracks[i].Likes)
		bLikes := len(tracks[j].Likes)

		if aLikes == bLikes {
			return tracks[i].Created > tracks[j].Created
		}

		return aLikes > bLikes
	})
}

// GetSoundTrack ...
func GetSoundTrack(id ID) (*SoundTrack, error) {
	track, err := DB.Get("SoundTrack", id)

	if err != nil {
		return nil, err
	}

	return track.(*SoundTrack), nil
}

// StreamSoundTracks returns a stream of all soundtracks.
func StreamSoundTracks() <-chan *SoundTrack {
	channel := make(chan *SoundTrack, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("SoundTrack") {
			channel <- obj.(*SoundTrack)
		}

		close(channel)
	}()

	return channel
}

// AllSoundTracks ...
func AllSoundTracks() []*SoundTrack {
	all := make([]*SoundTrack, 0, DB.Collection("SoundTrack").Count())

	for obj := range StreamSoundTracks() {
		all = append(all, obj)
	}

	return all
}

// FilterSoundTracks filters all soundtracks by a custom function.
func FilterSoundTracks(filter func(*SoundTrack) bool) []*SoundTrack {
	var filtered []*SoundTrack

	for obj := range StreamSoundTracks() {
		if filter(obj) {
			filtered = append(filtered, obj)
		}
	}

	return filtered
}
