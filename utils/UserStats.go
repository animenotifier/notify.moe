package utils

import (
	"time"

	"github.com/animenotifier/notify.moe/arn"
)

// UserStats ...
type UserStats struct {
	AnimeWatchingTime time.Duration
	PieCharts         []*arn.PieChart
}
