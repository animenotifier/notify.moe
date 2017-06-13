package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

const (
	networkRateLimit         = 100 * time.Millisecond
	avatarsDirectoryOriginal = "images/avatars/large/original/"
	avatarsDirectoryWebP     = "images/avatars/large/webp/"
)

var avatarSources []AvatarSource
var avatarWriters []AvatarWriter

// Main
func main() {
	// Switch to main directory
	os.Chdir("../../")

	// Define the avatar sources
	avatarSources = []AvatarSource{
		&Gravatar{},
	}

	// Stream of all users
	users, _ := arn.AllUsers()

	// Worker queue
	usersQueue := make(chan *arn.User)
	StartWorkers(usersQueue, networkRateLimit, Work)

	// We'll send each user to one of the worker threads
	for user := range users {
		usersQueue <- user
	}
}

// StartWorkers creates multiple workers to handle a user each.
func StartWorkers(queue chan *arn.User, rateLimit time.Duration, work func(*arn.User)) {
	rateLimiter := time.NewTicker(rateLimit)

	for w := 0; w < runtime.NumCPU(); w++ {
		go func() {
			for user := range queue {
				<-rateLimiter.C
				work(user)
			}
		}()
	}
}

// Work handles a single user.
func Work(user *arn.User) {
	for _, source := range avatarSources {
		avatar := source.GetAvatar(user)

		if avatar == nil {
			fmt.Println(color.RedString("✘"), user.Nick)
			continue
		}

		fmt.Println(color.GreenString("✔"), user.Nick, "|", avatar.Format, avatar.Image.Bounds().Dx(), avatar.Image.Bounds().Dy())

		for _, writer := range avatarWriters {
			writer.SaveAvatar(avatar)
		}

		return
	}

	// if downloadAvatar(user) {
	// 	makeWebPAvatar(user)
	// 	user.Avatar = "/+" + user.Nick + "/avatar"
	// } else {
	// 	user.Avatar = ""
	// }

	// user.Save()
}
