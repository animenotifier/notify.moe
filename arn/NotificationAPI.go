package arn

// Save saves the notification in the database.
func (notification *Notification) Save() {
	DB.Set("Notification", notification.ID, notification)
}
