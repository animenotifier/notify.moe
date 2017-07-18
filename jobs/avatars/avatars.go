package main

import (
	"os"
	"path"
	"runtime"
	"sync"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/aerogo/log"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/avatar/lib"
	"github.com/fatih/color"
)

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
	lib.Log.AddOutput(log.File("logs/avatar.log"))
	defer lib.Log.Flush()

	if InvokeShellArgs() {
		return
	}

	// Worker queue
	usersQueue := make(chan *arn.User, runtime.NumCPU())
	StartWorkers(usersQueue, lib.RefreshAvatar)

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
