package arn

// HasText includes a text field.
type hasText struct {
	Text string `json:"text" editable:"true" type:"textarea"`
}

// GetText returns the text of the object.
func (obj *hasText) GetText() string {
	return obj.Text
}
