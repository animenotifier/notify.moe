package arn

// Image represents an image with meta data.
type Image struct {
	Extension    string   `json:"extension"`
	Width        int      `json:"width"`
	Height       int      `json:"height"`
	AverageColor HSLColor `json:"averageColor"`
	LastModified int64    `json:"lastModified"`
}
