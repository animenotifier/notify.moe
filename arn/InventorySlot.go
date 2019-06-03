package arn

import "errors"

// InventorySlot ...
type InventorySlot struct {
	ItemID   string `json:"itemId"`
	Quantity uint   `json:"quantity"`
}

// IsEmpty ...
func (slot *InventorySlot) IsEmpty() bool {
	return slot.ItemID == ""
}

// Item ...
func (slot *InventorySlot) Item() *ShopItem {
	if slot.ItemID == "" {
		return nil
	}

	item, _ := GetShopItem(slot.ItemID)
	return item
}

// Decrease reduces the quantity by the given number.
func (slot *InventorySlot) Decrease(count uint) error {
	if slot.Quantity < count {
		return errors.New("Not enough items")
	}

	slot.Quantity -= count

	if slot.Quantity == 0 {
		slot.ItemID = ""
	}

	return nil
}

// Increase increases the quantity by the given number.
func (slot *InventorySlot) Increase(count uint) {
	slot.Quantity += count
}
