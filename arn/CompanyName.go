package arn

// CompanyName ...
type CompanyName struct {
	English  string   `json:"english" editable:"true"`
	Japanese string   `json:"japanese" editable:"true"`
	Synonyms []string `json:"synonyms" editable:"true"`
}
