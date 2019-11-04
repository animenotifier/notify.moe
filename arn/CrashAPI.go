package arn

import "github.com/aerogo/api"

// Force interface implementations
var (
	_ api.Savable = (*Crash)(nil)
)

// Save saves the crash in the database.
func (crash *Crash) Save() {
	DB.Set("Crash", crash.ID, crash)
}
