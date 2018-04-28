package jobs

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

var jobInfo = map[string]*utils.JobInfo{
	"anime-ratings": &utils.JobInfo{
		Name: "anime-ratings",
	},
	"twist": &utils.JobInfo{
		Name: "twist",
	},
	"refresh-osu": &utils.JobInfo{
		Name: "refresh-osu",
	},
	"mal-download": &utils.JobInfo{
		Name: "mal-download",
	},
	"mal-parse": &utils.JobInfo{
		Name: "mal-parse",
	},
	// "mal-sync": &utils.JobInfo{
	// 	Name: "mal-sync",
	// },
}

var jobLogs = []string{}

// Overview shows all background jobs.
func Overview(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	jobs := []*utils.JobInfo{}

	for _, job := range jobInfo {
		jobs = append(jobs, job)
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].Name < jobs[j].Name
	})

	return ctx.HTML(components.EditorJobs(jobs, jobLogs, ctx.URI(), user))
}
