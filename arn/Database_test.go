package arn_test

import (
	"testing"

	"github.com/animenotifier/notify.moe/arn"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	assert.NotEmpty(t, arn.DB.Node().Address().String())
}
