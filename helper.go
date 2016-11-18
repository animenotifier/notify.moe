package main

import "fmt"

// Converts anything into a string
func toString(v interface{}) string {
	return fmt.Sprint(v)
}

// Contains ...
func Contains(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Map ...
func Map(original []string, f func(string) string) []string {
	mapped := make([]string, len(original))
	for index, value := range original {
		mapped[index] = f(value)
	}
	return mapped
}
