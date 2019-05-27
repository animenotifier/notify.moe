package utils

import "github.com/akyoto/hash"

// HashStringsNoOrder returns a hash of the string slice contents, ignoring order.
func HashStringsNoOrder(items []string) uint64 {
	sum := uint64(0)

	for _, item := range items {
		sum += hash.String(item)
	}

	return sum
}
