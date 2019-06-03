package arn

import (
	"github.com/aerogo/api"
)

// Force interface implementations
var (
	_ Activity    = (*ActivityCreate)(nil)
	_ api.Savable = (*ActivityCreate)(nil)
)

// Save saves the activity object in the database.
func (activity *ActivityCreate) Save() {
	DB.Set("ActivityCreate", activity.ID, activity)
}

// Delete deletes the activity object from the database.
func (activity *ActivityCreate) Delete() error {
	DB.Delete("ActivityCreate", activity.ID)
	return nil
}
