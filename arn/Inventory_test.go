package arn_test

import (
	"testing"

	"github.com/akyoto/assert"

	"github.com/animenotifier/notify.moe/arn"
)

func TestInventory(t *testing.T) {
	inventory := arn.NewInventory("4J6qpK1ve")
	assert.Equal(t, len(inventory.Slots), arn.DefaultInventorySlotCount)
	assert.False(t, inventory.ContainsItem("pro-account-3"))

	err := inventory.AddItem("pro-account-3", 1)
	assert.Nil(t, err)
	assert.True(t, inventory.ContainsItem("pro-account-3"))
}
