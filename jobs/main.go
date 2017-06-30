package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/aerogo/log"
	"github.com/fatih/color"
)

var colorPool = []*color.Color{
	color.New(color.FgHiGreen),
	color.New(color.FgHiYellow),
	color.New(color.FgHiBlue),
	color.New(color.FgHiMagenta),
	color.New(color.FgHiCyan),
	color.New(color.FgGreen),
}

var jobs = map[string]time.Duration{
	"active-users":         1 * time.Minute,
	"forum-activity":       1 * time.Minute,
	"anime-ratings":        15 * time.Minute,
	"avatars":              30 * time.Minute,
	"refresh-track-titles": 10 * time.Hour,
	"sync-anime":           12 * time.Hour,
	"popular-anime":        12 * time.Hour,
	"airing-anime":         12 * time.Hour,
	"search-index":         12 * time.Hour,
}

func main() {
	// Start all jobs defined in the map above
	startJobs()

	// Wait for program termination
	wait()
}

func startJobs() {
	// Get the directory the executable is in
	exe, err := os.Executable()

	if err != nil {
		panic(err)
	}

	root := path.Dir(exe)

	// Log paths
	logsPath := path.Join(root, "../", "logs")
	jobLogsPath := path.Join(root, "../", "logs", "jobs")
	os.Mkdir(jobLogsPath, 0777)

	// Scheduler log
	mainLog := log.New()
	mainLog.AddOutput(os.Stdout)
	mainLog.AddOutput(log.File(path.Join(logsPath, "scheduler.log")))
	schedulerLog := mainLog

	// Color index
	colorIndex := 0

	// Start each job
	for job, interval := range jobs {
		jobName := job
		jobInterval := interval
		executable := path.Join(root, jobName, jobName)
		jobColor := colorPool[colorIndex].SprintFunc()

		jobLog := log.New()
		jobLog.AddOutput(log.File(path.Join(jobLogsPath, jobName+".log")))

		fmt.Printf("Registered job %s for execution every %v\n", jobColor(jobName), interval)

		go func() {
			ticker := time.NewTicker(jobInterval)
			defer ticker.Stop()

			var err error

			for {
				// Wait for the given interval first
				<-ticker.C

				// Now start
				schedulerLog.Info("Starting " + jobColor(jobName))

				cmd := exec.Command(executable)
				cmd.Stdout = jobLog
				cmd.Stderr = jobLog

				err = cmd.Start()

				if err != nil {
					schedulerLog.Error("Error starting job", jobColor(jobName), err)
				}

				err = cmd.Wait()

				if err != nil {
					schedulerLog.Error("Job exited with error", jobColor(jobName), err)
				}

				schedulerLog.Info("Finished " + jobColor(jobName))
				jobLog.Info("--------------------------------------------------------------------------------")
			}
		}()

		colorIndex = (colorIndex + 1) % len(colorPool)
	}

	// Finished job registration
	println("--------------------------------------------------------------------------------")
}

func wait() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}
