package utils

import (
	"github.com/aerogo/aero"
)

// GetContainerClass returns the class for the "container" element.
// In the browser extension it will get the "embedded" class.
func GetContainerClass(ctx *aero.Context) string {
	if ctx.URI() == "/extension/embed" {
		return "embedded"
	}

	return ""
}
