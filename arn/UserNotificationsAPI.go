package arn

// Force interface implementations
var (
	_ IDCollection = (*UserNotifications)(nil)
)

// Save saves the notification list in the database.
func (list *UserNotifications) Save() {
	DB.Set("UserNotifications", list.UserID, list)
}
