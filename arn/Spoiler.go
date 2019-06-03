package arn

// Spoiler represents a text that can spoil a future event.
type Spoiler struct {
	Text string `json:"text" editable:"true" type:"textarea"`
}

// String returns the containing text.
func (spoiler *Spoiler) String() string {
	return spoiler.Text
}
