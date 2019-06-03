package arn

// PushNotification represents a push notification.
type PushNotification struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Icon    string `json:"icon"`
	Link    string `json:"link"`
	Type    string `json:"type"`
}
