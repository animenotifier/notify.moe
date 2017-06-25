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
	color.New(color.FgCyan),
	color.New(color.FgYellow),
	color.New(color.FgGreen),
	color.New(color.FgBlue),
	color.New(color.FgMagenta),
}

var jobs = map[string]time.Duration{
	"popular-anime": 5 * time.Second,
	"search-index":  15 * time.Second,
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
				schedulerLog.Info("Starting " + jobColor(jobName))

				cmd := exec.Command(executable)
				cmd.Stdout = jobLog
				cmd.Stderr = jobLog

				err = cmd.Start()

				if err != nil {
					color.Red(err.Error())
				}

				err = cmd.Wait()

				if err != nil {
					color.Red(err.Error())
				}

				schedulerLog.Info("Finished " + jobColor(jobName))
				jobLog.Info("--------------------------------------------------------------------------------")

				<-ticker.C
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
