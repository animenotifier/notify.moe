package arn_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/animenotifier/notify.moe/arn"
)

func TestNewUser(t *testing.T) {
	user := arn.NewUser()

	assert.NotNil(t, user)
	assert.NotEqual(t, user.ID, "")
}

func TestDatabaseErrorMessages(t *testing.T) {
	_, err := arn.GetUser("NON EXISTENT USER ID")

	// We need to make sure that non-existent records return "not found" in the error message.
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "not found")
}
