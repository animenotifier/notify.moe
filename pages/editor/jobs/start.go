package jobs

import (
	"net/http"
	"sync"

	"github.com/animenotifier/arn"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/utils"
)

// Only allow one job to be started at a time
var jobStartMutex sync.Mutex

// Start will start the specified background job.
func Start(ctx aero.Context) error {
	jobStartMutex.Lock()
	defer jobStartMutex.Unlock()

	user := utils.GetUser(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not authorized")
	}

	jobName := ctx.Get("job")
	job := jobInfo[jobName]

	if job == nil {
		return ctx.Error(http.StatusBadRequest, "Job not available")
	}

	if job.IsRunning() {
		return ctx.Error(http.StatusBadRequest, "Job is currently running!")
	}

	job.Start()
	jobLogs = append(jobLogs, user.Nick+" started "+job.Name+" job ("+arn.DateTimeUTC()+").")

	return nil
}
