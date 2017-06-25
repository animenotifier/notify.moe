package main

import (
	"fmt"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/fatih/color"
)

var jobs = map[string]time.Duration{
	"popular-anime": 1 * time.Second,
}

func main() {
	// Start all jobs defined in the map above
	startJobs()

	// Wait for program termination
	wait()
}

func startJobs() {
	exe, err := os.Executable()

	if err != nil {
		panic(err)
	}

	root := path.Dir(exe)

	for job, interval := range jobs {
		jobInterval := interval
		executable := path.Join(root, job, job)

		fmt.Printf("Registered job %s for execution every %v\n", color.YellowString(job), interval)

		go func() {
			ticker := time.NewTicker(jobInterval)
			defer ticker.Stop()

			for {
				fmt.Println(executable)
				<-ticker.C
			}
		}()
	}
}

func wait() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}
