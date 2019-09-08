package arn

import "github.com/aerogo/api"

// Force interface implementations
var (
	_ Identifiable = (*DraftIndex)(nil)
	_ api.Savable  = (*DraftIndex)(nil)
)

// Save saves the index in the database.
func (index *DraftIndex) Save() {
	DB.Set("DraftIndex", index.UserID, index)
}
