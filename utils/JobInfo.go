package utils

import "time"

// JobInfo gives you information about a background job.
type JobInfo struct {
	Name         string
	LastStarted  time.Time
	LastFinished time.Time
}

// IsRunning tells you whether the given job is running or not.
func (job *JobInfo) IsRunning() bool {
	now := time.Now()
	return job.LastStarted.After(job.LastFinished) && !now.After(job.LastFinished)
}
