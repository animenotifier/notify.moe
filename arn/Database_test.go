package arn_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/animenotifier/notify.moe/arn"
)

func TestConnect(t *testing.T) {
	assert.NotEqual(t, arn.DB.Node().Address().String(), "")
}
