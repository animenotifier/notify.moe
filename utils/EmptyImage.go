package utils

// EmptyImage returns the smallest possible 1x1 pixel image encoded in Base64.
func EmptyImage() string {
	return ""
	// return "data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw=="
}
