package main

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"runtime"
	"sync"
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
var wg sync.WaitGroup

// Main
func main() {
	color.Yellow("Generating user avatars")

	// Switch to main directory
	exe, err := os.Executable()

	if err != nil {
		panic(err)
	}

	root := path.Dir(exe)
	os.Chdir(path.Join(root, "../../"))

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
		&FileSystem{
			Directory: "images/avatars/large/",
		},
	}

	// Define the avatar outputs
	avatarOutputs = []AvatarOutput{
		// Original - Large
		&AvatarOriginalFileOutput{
			Directory: "images/avatars/large/",
			Size:      arn.AvatarMaxSize,
		},

		// Original - Small
		&AvatarOriginalFileOutput{
			Directory: "images/avatars/small/",
			Size:      arn.AvatarSmallSize,
		},

		// WebP - Large
		&AvatarWebPFileOutput{
			Directory: "images/avatars/large/",
			Size:      arn.AvatarMaxSize,
			Quality:   webPQuality,
		},

		// WebP - Small
		&AvatarWebPFileOutput{
			Directory: "images/avatars/small/",
			Size:      arn.AvatarSmallSize,
			Quality:   webPQuality,
		},
	}

	if InvokeShellArgs() {
		return
	}

	// Worker queue
	usersQueue := make(chan *arn.User, runtime.NumCPU())
	StartWorkers(usersQueue, Work)

	allUsers, _ := arn.AllUsers()

	// We'll send each user to one of the worker threads
	for _, user := range allUsers {
		wg.Add(1)
		usersQueue <- user
	}

	wg.Wait()

	color.Green("Finished.")
}

// StartWorkers creates multiple workers to handle a user each.
func StartWorkers(queue chan *arn.User, work func(*arn.User)) {
	for w := 0; w < runtime.NumCPU(); w++ {
		go func() {
			for user := range queue {
				work(user)
				wg.Done()
			}
		}()
	}
}

// Work handles a single user.
func Work(user *arn.User) {
	user.AvatarExtension = ""

	for _, source := range avatarSources {
		avatar := source.GetAvatar(user)

		if avatar == nil {
			// fmt.Println(color.RedString("✘"), reflect.TypeOf(source).Elem().Name(), user.Nick)
			continue
		}

		// Name of source
		sourceType := reflect.TypeOf(source).Elem().Name()

		// Log
		fmt.Println(color.GreenString("✔"), sourceType, "|", user.Nick, "|", avatar)

		// Avoid quality loss (if it's on the file system, we don't need to write it again)
		if sourceType == "FileSystem" {
			user.AvatarExtension = avatar.Extension()
			break
		}

		for _, writer := range avatarOutputs {
			err := writer.SaveAvatar(avatar)

			if err != nil {
				color.Red(err.Error())
			}
		}

		break
	}

	// Since this a very long running job, refresh user data before saving it.
	avatarExt := user.AvatarExtension
	user, err := arn.GetUser(user.ID)

	if err != nil {
		avatarLog.Error("Can't refresh user info:", user.ID, user.Nick)
		return
	}

	// Save avatar data
	user.AvatarExtension = avatarExt
	user.Save()
}
