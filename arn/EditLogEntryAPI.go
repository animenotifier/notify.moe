package arn

// Save saves the log entry in the database.
func (entry *EditLogEntry) Save() {
	DB.Set("EditLogEntry", entry.ID, entry)
}
