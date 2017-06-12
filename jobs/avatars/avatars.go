package main

import (
	"os"
	"runtime"
	"time"

	"github.com/animenotifier/arn"
)

const (
	avatarsDirectoryOriginal = "images/avatars/original/"
	avatarsDirectoryWebP     = "images/avatars/webp/"
)

func main() {
	os.Chdir("../../")

	users, _ := arn.AllUsers()

	usersQueue := make(chan *arn.User)
	rateLimiter := time.NewTicker(100 * time.Millisecond)
	defer rateLimiter.Stop()

	for w := 0; w < runtime.NumCPU(); w++ {
		go func(workerID int) {
			for user := range usersQueue {
				<-rateLimiter.C

				if downloadAvatar(user) {
					makeWebPAvatar(user)
					user.Avatar = "/+" + user.Nick + "/avatar"
				} else {
					user.Avatar = ""
				}

				user.Save()
			}
		}(w)
	}

	for user := range users {
		usersQueue <- user
	}
}
