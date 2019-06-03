package arn

import (
	"errors"
)

// DefaultInventorySlotCount tells you how many slots are available by default in an inventory.
const DefaultInventorySlotCount = 24

// Inventory has inventory slots that store shop item IDs and their quantity.
type Inventory struct {
	UserID UserID           `json:"userId"`
	Slots  []*InventorySlot `json:"slots"`
}

// AddItem adds a given item to the inventory.
func (inventory *Inventory) AddItem(itemID string, quantity uint) error {
	if itemID == "" {
		return nil
	}

	// Find the slot with the item
	for _, slot := range inventory.Slots {
		if slot.ItemID == itemID {
			slot.Quantity += quantity
			return nil
		}
	}

	// If the item doesn't exist in the inventory yet, add it to the first free slot
	for _, slot := range inventory.Slots {
		if slot.ItemID == "" {
			slot.ItemID = itemID
			slot.Quantity = quantity
			return nil
		}
	}

	// If there is no free slot, return an error
	return errors.New("Inventory is full")
}

// ContainsItem checks if the inventory contains the item ID already.
func (inventory *Inventory) ContainsItem(itemID string) bool {
	for _, slot := range inventory.Slots {
		if slot.ItemID == itemID {
			return true
		}
	}

	return false
}

// SwapSlots swaps the slots with the given indices.
func (inventory *Inventory) SwapSlots(a, b int) error {
	if a < 0 || b < 0 || a >= len(inventory.Slots) || b >= len(inventory.Slots) {
		return errors.New("Inventory slot index out of bounds")
	}

	// Swap
	inventory.Slots[a], inventory.Slots[b] = inventory.Slots[b], inventory.Slots[a]
	return nil
}

// NewInventory creates a new inventory with the default number of slots.
func NewInventory(userID UserID) *Inventory {
	inventory := &Inventory{
		UserID: userID,
		Slots:  make([]*InventorySlot, DefaultInventorySlotCount),
	}

	for i := 0; i < len(inventory.Slots); i++ {
		inventory.Slots[i] = &InventorySlot{}
	}

	return inventory
}

// GetInventory ...
func GetInventory(userID UserID) (*Inventory, error) {
	obj, err := DB.Get("Inventory", userID)

	if err != nil {
		return nil, err
	}

	return obj.(*Inventory), nil
}
