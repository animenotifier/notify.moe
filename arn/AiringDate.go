package arn

import (
	"time"
)

// AiringDate represents the airing date of an anime.
type AiringDate struct {
	Start string `json:"start" editable:"true"`
	End   string `json:"end" editable:"true"`
}

// StartDateHuman returns the start date of the anime in human readable form.
func (airing *AiringDate) StartDateHuman() string {
	t, _ := time.Parse(time.RFC3339, airing.Start)
	humanReadable := t.Format(time.RFC1123)

	return humanReadable[:len("Thu, 25 May 2017")]
}

// EndDateHuman returns the end date of the anime in human readable form.
func (airing *AiringDate) EndDateHuman() string {
	t, _ := time.Parse(time.RFC3339, airing.End)
	humanReadable := t.Format(time.RFC1123)

	return humanReadable[:len("Thu, 25 May 2017")]
}

// StartTimeHuman returns the start time of the anime in human readable form.
func (airing *AiringDate) StartTimeHuman() string {
	t, _ := time.Parse(time.RFC3339, airing.Start)
	humanReadable := t.Format(time.RFC1123)

	return humanReadable[len("Thu, 25 May 2017 "):]
}

// EndTimeHuman returns the end time of the anime in human readable form.
func (airing *AiringDate) EndTimeHuman() string {
	t, _ := time.Parse(time.RFC3339, airing.End)
	humanReadable := t.Format(time.RFC1123)

	return humanReadable[len("Thu, 25 May 2017 "):]
}
