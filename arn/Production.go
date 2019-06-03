package arn

import (
	"os"
	"strings"
)

// IsProduction returns true if the hostname contains "arn".
func IsProduction() bool {
	return strings.Contains(HostName(), "arn")
}

// IsDevelopment returns true if the hostname does not contain "arn".
func IsDevelopment() bool {
	return !IsProduction()
}

// HostName returns the hostname.
func HostName() string {
	host, _ := os.Hostname()
	return host
}
