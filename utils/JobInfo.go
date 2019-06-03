package utils

import (
	"os/exec"
	"path"
	"time"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

// JobInfo gives you information about a background job.
type JobInfo struct {
	Name         string
	LastStarted  time.Time
	LastFinished time.Time
}

// IsRunning tells you whether the given job is running or not.
func (job *JobInfo) IsRunning() bool {
	return job.LastStarted.After(job.LastFinished)
}

// Start will start the job.
func (job *JobInfo) Start() error {
	cmd := exec.Command(path.Join(arn.Root, "jobs", job.Name, job.Name))
	err := cmd.Start()

	if err != nil {
		return err
	}

	job.LastStarted = time.Now()

	// Wait for job finish in another goroutine
	go func() {
		err := cmd.Wait()

		if err != nil {
			color.Red("Job '%s' encountered an error: %s", job.Name, err.Error())
		}

		job.LastFinished = time.Now()
	}()

	return nil
}
