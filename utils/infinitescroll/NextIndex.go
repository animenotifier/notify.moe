package infinitescroll

import (
	"strconv"

	"github.com/aerogo/aero"
)

// NextIndex calculates the next index and sends HTTP header
func NextIndex(ctx aero.Context, allElementsLength int, elementsPerScroll int, index int) int {
	nextIndex := index + elementsPerScroll

	if nextIndex >= allElementsLength {
		nextIndex = -1
	}

	// Send the index for the next request
	ctx.Response().SetHeader("X-LoadMore-Index", strconv.Itoa(nextIndex))

	return nextIndex
}
