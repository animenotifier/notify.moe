package arn

// Save saves the item in the database.
func (item *ShopItem) Save() {
	DB.Set("ShopItem", item.ID, item)
}
