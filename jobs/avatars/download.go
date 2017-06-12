package main

import (
	"fmt"
	"io/ioutil"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
	"github.com/parnurzeal/gorequest"
	gravatar "github.com/ungerik/go-gravatar"
)

func findOriginalAvatar(user *arn.User) string {
	return arn.FindFileWithExtension(
		user.ID,
		avatarsDirectoryOriginal,
		[]string{
			".jpg",
			".png",
			".gif",
		},
	)
}

func downloadAvatar(user *arn.User) bool {
	if user.Email == "" {
		return false
	}

	directory := avatarsDirectoryOriginal
	fileName := directory + user.ID

	// Build URL
	url := gravatar.Url(user.Email) + "?s=" + fmt.Sprint(arn.AvatarMaxSize) + "&d=404&r=pg"

	// Skip existing avatars
	if findOriginalAvatar(user) != "" {
		color.Yellow(user.Nick)
		return true
	}

	// Download
	response, data, err := gorequest.New().Get(url).EndBytes()

	if err != nil {
		color.Red(user.Nick)
		return false
	}

	contentType := response.Header.Get("content-type")

	if response.StatusCode != 200 {
		color.Red(user.Nick)
		return false
	}

	// Determine file extension
	extension := ""

	switch contentType {
	case "image/jpeg":
		extension = ".jpg"
	case "image/png":
		extension = ".png"
	case "image/gif":
		extension = ".gif"
	case "image/webp":
		extension = ".webp"
	}

	fileName += extension

	// Write to disk
	ioutil.WriteFile(fileName, data, 0644)

	color.Green(user.Nick)

	return true
}
