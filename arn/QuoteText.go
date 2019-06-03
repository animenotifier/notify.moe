package arn

// QuoteText ...
type QuoteText struct {
	English  string `json:"english" editable:"true" type:"textarea"`
	Japanese string `json:"japanese" editable:"true" type:"textarea"`
}
