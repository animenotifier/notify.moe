package utils

import "time"
import "github.com/animenotifier/arn"

// UserStats ...
type UserStats struct {
	AnimeWatchingTime time.Duration
	PieCharts         []*arn.PieChart
}
