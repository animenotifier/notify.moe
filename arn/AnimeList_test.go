package arn_test

import (
	"testing"

	"github.com/animenotifier/notify.moe/arn"
)

func TestNormalizeRatings(t *testing.T) {
	user, _ := arn.GetUser("4J6qpK1ve")
	animeList := user.AnimeList()
	animeList.NormalizeRatings()
}
