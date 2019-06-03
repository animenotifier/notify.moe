package arn

// Link describes a single link to an external website.
type Link struct {
	Title string `json:"title" editable:"true"`
	URL   string `json:"url" editable:"true"`
}
