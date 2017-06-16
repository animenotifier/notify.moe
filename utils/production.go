package utils

import (
	"os"
	"strings"
)

// IsProduction returns true if the hostname contains "arn".
func IsProduction() bool {
	host, _ := os.Hostname()
	return strings.Contains(host, "arn")
}

// IsDevelopment returns true if the hostname does not contain "arn".
func IsDevelopment() bool {
	return !IsProduction()
}
