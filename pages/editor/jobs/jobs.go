package jobs

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

var jobInfo = map[string]*utils.JobInfo{
	"kitsu-import-anime": {
		Name: "kitsu-import-anime",
	},
	// "anime-ratings": &utils.JobInfo{
	// 	Name: "anime-ratings",
	// },
	// "twist": &utils.JobInfo{
	// 	Name: "twist",
	// },
	// "refresh-games": &utils.JobInfo{
	// 	Name: "refresh-games",
	// },
	// "test": &utils.JobInfo{
	// 	Name: "test",
	// },
}

var jobLogs = []string{}

// Overview shows all background jobs.
func Overview(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	jobs := []*utils.JobInfo{}

	for _, job := range jobInfo {
		jobs = append(jobs, job)
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].Name < jobs[j].Name
	})

	return ctx.HTML(components.EditorJobs(jobs, jobLogs, ctx.Path(), user))
}
