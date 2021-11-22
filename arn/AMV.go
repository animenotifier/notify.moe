package arn

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/aerogo/nano"
	"github.com/animenotifier/notify.moe/arn/video"
	"github.com/minio/minio-go/v7"
)

// AMV is an anime music video.
type AMV struct {
	File           string     `json:"file" editable:"true" type:"upload" filetype:"video" endpoint:"/api/upload/amv/:id/file"`
	Title          AMVTitle   `json:"title" editable:"true"`
	MainAnimeID    AnimeID    `json:"mainAnimeId" editable:"true"`
	ExtraAnimeIDs  []AnimeID  `json:"extraAnimeIds" editable:"true"`
	VideoEditorIDs []UserID   `json:"videoEditorIds" editable:"true"`
	Links          []Link     `json:"links" editable:"true"`
	Tags           []string   `json:"tags" editable:"true"`
	Info           video.Info `json:"info"`

	hasID
	hasPosts
	hasCreator
	hasEditor
	hasLikes
	hasDraft
}

// Link returns the permalink for the AMV.
func (amv *AMV) Link() string {
	return "/amv/" + amv.ID
}

// VideoLink returns the permalink for the video file.
func (amv *AMV) VideoLink() string {
	domain := "arn.sfo2.cdn"

	if amv.IsDraft {
		domain = "arn.sfo2"
	}

	return fmt.Sprintf("https://%s.digitaloceanspaces.com/videos/amvs/%s", domain, amv.File)
}

// TitleByUser returns the preferred title for the given user.
func (amv *AMV) TitleByUser(user *User) string {
	return amv.Title.ByUser(user)
}

// SetVideoReader sets the bytes for the video file by reading them from the reader.
func (amv *AMV) SetVideoReader(reader io.Reader) error {
	fileName := amv.ID + ".webm"
	pattern := amv.ID + ".*.webm"
	file, err := ioutil.TempFile("", pattern)

	if err != nil {
		return err
	}

	filePath := file.Name()
	defer os.Remove(filePath)

	// Write file contents
	_, err = io.Copy(file, reader)

	if err != nil {
		return err
	}

	// Run mkclean
	optimizedFile := filePath + ".optimized"
	defer os.Remove(optimizedFile)

	cmd := exec.Command(
		"mkclean",
		"--doctype", "4",
		"--keep-cues",
		"--optimize",
		filePath,
		optimizedFile,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Start()

	if err != nil {
		return err
	}

	err = cmd.Wait()

	if err != nil {
		return err
	}

	// Refresh video file info
	info, err := video.GetInfo(optimizedFile)

	if err != nil {
		return err
	}

	// Is our storage server available?
	if Spaces == nil {
		return errors.New("File storage client has not been initialized")
	}

	// Make sure the file is public
	userMetaData := map[string]string{
		"x-amz-acl": "public-read",
	}

	// Upload the file to our storage server
	_, err = Spaces.FPutObject(context.TODO(), "arn", fmt.Sprintf("videos/amvs/%s.webm", amv.ID), optimizedFile, minio.PutObjectOptions{
		ContentType:  "video/webm",
		UserMetadata: userMetaData,
	})

	if err != nil {
		return err
	}

	amv.Info = *info
	amv.File = fileName
	return nil
}

// MainAnime returns main anime for the AMV, if available.
func (amv *AMV) MainAnime() *Anime {
	mainAnime, _ := GetAnime(amv.MainAnimeID)
	return mainAnime
}

// ExtraAnime returns main anime for the AMV, if available.
func (amv *AMV) ExtraAnime() []*Anime {
	objects := DB.GetMany("Anime", amv.ExtraAnimeIDs)
	animes := make([]*Anime, 0, len(amv.ExtraAnimeIDs))

	for _, obj := range objects {
		if obj == nil {
			continue
		}

		animes = append(animes, obj.(*Anime))
	}

	return animes
}

// VideoEditors returns a slice of all the users involved in creating the AMV.
func (amv *AMV) VideoEditors() []*User {
	objects := DB.GetMany("User", amv.VideoEditorIDs)
	editors := []*User{}

	for _, obj := range objects {
		if obj == nil {
			continue
		}

		editors = append(editors, obj.(*User))
	}

	return editors
}

// Publish turns the draft into a published object.
func (amv *AMV) Publish() error {
	// No title
	if amv.Title.String() == "" {
		return errors.New("AMV doesn't have a title")
	}

	// No anime found
	if amv.MainAnimeID == "" && len(amv.ExtraAnimeIDs) == 0 {
		return errors.New("Need to specify at least one anime")
	}

	// Check that the file name exists
	if amv.File == "" {
		return errors.New("You need to upload a WebM file for this AMV")
	}

	// No file uploaded
	_, err := Spaces.StatObject(context.TODO(), "arn", fmt.Sprintf("videos/amvs/%s", amv.File), minio.StatObjectOptions{})

	if err != nil {
		return errors.New("You need to upload a WebM file for this AMV")
	}

	return publish(amv)
}

// Unpublish turns the object back into a draft.
func (amv *AMV) Unpublish() error {
	return unpublish(amv)
}

// OnLike is called when the AMV receives a like.
func (amv *AMV) OnLike(likedBy *User) {
	if likedBy.ID == amv.CreatedBy {
		return
	}

	go func() {
		amv.Creator().SendNotification(&PushNotification{
			Title:   likedBy.Nick + " liked your AMV " + amv.Title.ByUser(amv.Creator()),
			Message: likedBy.Nick + " liked your AMV " + amv.Title.ByUser(amv.Creator()) + ".",
			Icon:    "https:" + likedBy.AvatarLink("large"),
			Link:    "https://notify.moe" + likedBy.Link(),
			Type:    NotificationTypeLike,
		})
	}()
}

// String implements the default string serialization.
func (amv *AMV) String() string {
	return amv.Title.ByUser(nil)
}

// TypeName returns the type name.
func (amv *AMV) TypeName() string {
	return "AMV"
}

// Self returns the object itself.
func (amv *AMV) Self() Loggable {
	return amv
}

// GetAMV returns the AMV with the given ID.
func GetAMV(id ID) (*AMV, error) {
	obj, err := DB.Get("AMV", id)

	if err != nil {
		return nil, err
	}

	return obj.(*AMV), nil
}

// StreamAMVs returns a stream of all AMVs.
func StreamAMVs() <-chan *AMV {
	channel := make(chan *AMV, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("AMV") {
			channel <- obj.(*AMV)
		}

		close(channel)
	}()

	return channel
}

// AllAMVs returns a slice of all AMVs.
func AllAMVs() []*AMV {
	all := make([]*AMV, 0, DB.Collection("AMV").Count())

	stream := StreamAMVs()

	for obj := range stream {
		all = append(all, obj)
	}

	return all
}

// FilterAMVs filters all AMVs by a custom function.
func FilterAMVs(filter func(*AMV) bool) []*AMV {
	var filtered []*AMV

	for obj := range DB.All("AMV") {
		realObject := obj.(*AMV)

		if filter(realObject) {
			filtered = append(filtered, realObject)
		}
	}

	return filtered
}
