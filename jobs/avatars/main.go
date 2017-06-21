package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/aerogo/log"
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

const (
	webPQuality = 80
)

var avatarSources []AvatarSource
var avatarOutputs []AvatarOutput
var avatarLog = log.New()

// Main
func main() {
	color.Yellow("Generating user avatars")

	// Switch to main directory
	os.Chdir("../../")

	// Log
	avatarLog.AddOutput(log.File("logs/avatar.log"))
	defer avatarLog.Flush()

	// Define the avatar sources
	avatarSources = []AvatarSource{
		&Gravatar{
			Rating:         "pg",
			RequestLimiter: time.NewTicker(250 * time.Millisecond),
		},
		&MyAnimeList{
			RequestLimiter: time.NewTicker(250 * time.Millisecond),
		},
	}

	// Define the avatar outputs
	avatarOutputs = []AvatarOutput{
		// Original - Large
		&AvatarOriginalFileOutput{
			Directory: "images/avatars/large/original/",
			Size:      arn.AvatarMaxSize,
		},

		// Original - Small
		&AvatarOriginalFileOutput{
			Directory: "images/avatars/small/original/",
			Size:      arn.AvatarSmallSize,
		},

		// WebP - Large
		&AvatarWebPFileOutput{
			Directory: "images/avatars/large/webp/",
			Size:      arn.AvatarMaxSize,
			Quality:   webPQuality,
		},

		// WebP - Small
		&AvatarWebPFileOutput{
			Directory: "images/avatars/small/webp/",
			Size:      arn.AvatarSmallSize,
			Quality:   webPQuality,
		},
	}

	// Stream of all users
	users, _ := arn.FilterUsers(func(user *arn.User) bool {
		return true
	})

	// Log user count
	println(len(users), "users")

	// Worker queue
	usersQueue := make(chan *arn.User)
	StartWorkers(usersQueue, Work)

	// We'll send each user to one of the worker threads
	for _, user := range users {
		usersQueue <- user
	}

	color.Green("Finished.")
}

// StartWorkers creates multiple workers to handle a user each.
func StartWorkers(queue chan *arn.User, work func(*arn.User)) {
	for w := 0; w < runtime.NumCPU(); w++ {
		go func() {
			for user := range queue {
				work(user)
			}
		}()
	}
}

// Work handles a single user.
func Work(user *arn.User) {
	user.Avatar = ""

	for _, source := range avatarSources {
		avatar := source.GetAvatar(user)

		if avatar == nil {
			// fmt.Println(color.RedString("✘"), reflect.TypeOf(source).Elem().Name(), user.Nick)
			continue
		}

		for _, writer := range avatarOutputs {
			err := writer.SaveAvatar(avatar)

			if err != nil {
				color.Red(err.Error())
			}
		}

		fmt.Println(color.GreenString("✔"), reflect.TypeOf(source).Elem().Name(), "|", user.Nick, "|", avatar)
		user.Avatar = "/+" + user.Nick + "/avatar"
		break
	}

	// Save avatar data
	user.Save()
}
