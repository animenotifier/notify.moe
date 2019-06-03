package arn

// Save saves the purchase in the database.
func (purchase *Purchase) Save() {
	DB.Set("Purchase", purchase.ID, purchase)
}
