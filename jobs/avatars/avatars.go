package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"time"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
	"github.com/parnurzeal/gorequest"
	gravatar "github.com/ungerik/go-gravatar"
)

func main() {
	users, _ := arn.AllUsers()

	usersQueue := make(chan *arn.User)
	rateLimiter := time.NewTicker(100 * time.Millisecond)
	defer rateLimiter.Stop()

	for w := 0; w < runtime.NumCPU(); w++ {
		go func(workerID int) {
			for user := range usersQueue {
				<-rateLimiter.C
				os.Stdout.WriteString("[" + fmt.Sprint(workerID) + "] ")
				downloadAvatar(user)
				makeWebPAvatar(user)
			}
		}(w)
	}

	for user := range users {
		usersQueue <- user
	}
}

func findAvatar(user *arn.User, dir string) string {
	testExtensions := []string{"", ".jpg", ".png", ".gif", ".webp"}

	for _, testExt := range testExtensions {
		if _, err := os.Stat(dir + user.ID + testExt); !os.IsNotExist(err) {
			return user.ID + testExt
		}
	}

	return ""
}

func makeWebPAvatar(user *arn.User) {
	baseName := findAvatar(user, "../../images/avatars/original/")

	if baseName == "" {
		return
	}

	original := "../../images/avatars/original/" + baseName
	outFile := "../../images/avatars/webp/" + user.ID + ".webp"

	err := convertFileToWebP(original, outFile, 80)

	if err != nil {
		color.Red("[WebP] " + original + " -> " + outFile)
	} else {
		color.Green("[WebP] " + original + " -> " + outFile)
	}
}

func downloadAvatar(user *arn.User) {
	if user.Email == "" {
		return
	}

	directory := "../../images/avatars/original/"
	fileName := directory + user.ID

	// Build URL
	url := gravatar.Url(user.Email) + "?s=560&d=404&r=pg"

	// Skip existing avatars
	if findAvatar(user, directory) != "" {
		color.Yellow(url)
		return
	}

	// Download
	response, data, err := gorequest.New().Get(url).EndBytes()

	if err != nil {
		color.Red(url)
		return
	}

	contentType := response.Header.Get("content-type")

	if response.StatusCode != 200 {
		color.Red(url)
		return
	}

	color.Green(url)

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
}
