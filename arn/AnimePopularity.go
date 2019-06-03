package arn

// AnimePopularity shows how many users have that anime in a certain list.
type AnimePopularity struct {
	Watching  int `json:"watching"`
	Completed int `json:"completed"`
	Planned   int `json:"planned"`
	Hold      int `json:"hold"`
	Dropped   int `json:"dropped"`
}

// Total returns the total number of users that added this anime to their collection.
func (p *AnimePopularity) Total() int {
	return p.Watching + p.Completed + p.Planned + p.Hold + p.Dropped
}
