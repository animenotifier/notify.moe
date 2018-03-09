package utils

import "github.com/OneOfOne/xxhash"

// HashString returns a hash of the string.
func HashString(item string) uint64 {
	h := xxhash.NewS64(0)
	h.Write([]byte(item))
	return h.Sum64()
}

// HashStringsNoOrder returns a hash of the string slice contents, ignoring order.
func HashStringsNoOrder(items []string) uint64 {
	sum := uint64(0)

	for _, item := range items {
		h := xxhash.NewS64(0)
		h.Write([]byte(item))
		numHash := h.Sum64()
		sum += numHash
	}

	return sum
}
