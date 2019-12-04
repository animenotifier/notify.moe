package arn

import (
	"os"
)

// IsProduction returns true if PRODUCTION is set to 1.
func IsProduction() bool {
	return os.Getenv("PRODUCTION") == "1"
}

// IsDevelopment returns true if PRODUCTION is not set to 1.
func IsDevelopment() bool {
	return !IsProduction()
}
