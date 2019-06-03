package arn

import (
	"errors"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

// Actions
func init() {
	API.RegisterActions("Inventory", []*api.Action{
		// Use slot
		{
			Name:  "use",
			Route: "/use/:slot",
			Run: func(obj interface{}, ctx aero.Context) error {
				inventory := obj.(*Inventory)
				slotIndex, err := ctx.GetInt("slot")

				if err != nil {
					return err
				}

				slot := inventory.Slots[slotIndex]

				if slot.IsEmpty() {
					return errors.New("No item in this slot")
				}

				if !slot.Item().Consumable {
					return errors.New("This item is not consumable")
				}

				// Save item ID in case it gets deleted by slot.Decrease()
				itemID := slot.ItemID

				// Decrease quantity
				err = slot.Decrease(1)

				if err != nil {
					return err
				}

				// Save inventory
				inventory.Save()

				user := GetUserFromContext(ctx)
				err = user.ActivateItemEffect(itemID)

				if err != nil {
					// Refund item
					slot.ItemID = itemID
					slot.Increase(1)
					inventory.Save()
					return nil
				}

				return err
			},
		},

		// Swap slots
		{
			Name:  "swap",
			Route: "/swap/:slot1/:slot2",
			Run: func(obj interface{}, ctx aero.Context) error {
				inventory := obj.(*Inventory)
				a, err := ctx.GetInt("slot1")

				if err != nil {
					return err
				}

				b, err := ctx.GetInt("slot2")

				if err != nil {
					return err
				}

				err = inventory.SwapSlots(a, b)

				if err != nil {
					return err
				}

				inventory.Save()

				return nil
			},
		},
	})
}

// Authorize returns an error if the given API request is not authorized.
func (inventory *Inventory) Authorize(ctx aero.Context, action string) error {
	return AuthorizeIfLoggedInAndOwnData(ctx, "id")
}

// Save saves the push items in the database.
func (inventory *Inventory) Save() {
	DB.Set("Inventory", inventory.UserID, inventory)
}
