package arn

// Feature represents a donation-based feature on the site.
type Feature struct {
	RequiredAmount int `json:"requiredAmount" editable:"true"`
	ReceivedAmount int `json:"receivedAmount"`

	hasID
}
