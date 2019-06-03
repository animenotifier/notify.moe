package arn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/animenotifier/notify.moe/arn"
)

func TestInventory(t *testing.T) {
	inventory := arn.NewInventory("4J6qpK1ve")

	assert.Len(t, inventory.Slots, arn.DefaultInventorySlotCount)
	assert.False(t, inventory.ContainsItem("pro-account-3"))

	err := inventory.AddItem("pro-account-3", 1)
	assert.NoError(t, err)
	assert.True(t, inventory.ContainsItem("pro-account-3"))
}
