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
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
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
	"anime-ratings": 10 * time.Minute,
	"twist":         2 * time.Hour,
	"refresh-games": 6 * time.Hour,
}

func main() {
	// Start all jobs defined in the map above
	startJobs()

	// Wait for program termination
	wait()
}

func startJobs() {
	// Log paths
	logsPath := path.Join(arn.Root, "logs")
	jobLogsPath := path.Join(arn.Root, "logs", "jobs")
	_ = os.Mkdir(jobLogsPath, 0777)

	// Scheduler log
	mainLog := log.New()
	mainLog.AddWriter(os.Stdout)
	mainLog.AddWriter(log.File(path.Join(logsPath, "scheduler.log")))
	schedulerLog := mainLog

	// Color index
	colorIndex := 0

	// Start each job
	for job, interval := range jobs {
		jobName := job
		jobInterval := interval
		executable := path.Join(arn.Root, "jobs", jobName, jobName)
		jobColor := colorPool[colorIndex].SprintFunc()

		jobLog := log.New()
		jobLog.AddWriter(log.File(path.Join(jobLogsPath, jobName+".log")))

		fmt.Printf("Registered job %s for execution every %v\n", jobColor(jobName), interval)

		go func() {
			ticker := time.NewTicker(jobInterval)
			defer ticker.Stop()

			for {
				// Wait for the given interval first
				<-ticker.C

				// Now start
				schedulerLog.Info("Starting %s", jobColor(jobName))

				cmd := exec.Command(executable)
				cmd.Stdout = jobLog
				cmd.Stderr = jobLog

				err := cmd.Start()

				if err != nil {
					schedulerLog.Error("Error starting job %s %v", jobColor(jobName), err)
				}

				err = cmd.Wait()

				if err != nil {
					schedulerLog.Error("Job %s exited with error %v", jobColor(jobName), err)
				}

				schedulerLog.Info("Finished %s", jobColor(jobName))
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
