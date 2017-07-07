package utils

import (
	"fmt"
	"math"
)

// coords returns the coordinates for the given percentage.
func coords(percent float64) (float64, float64) {
	x := math.Cos(2 * math.Pi * percent)
	y := math.Sin(2 * math.Pi * percent)
	return x, y
}

// SVGSlicePath creates a path string for a slice in a pie chart.
func SVGSlicePath(from float64, to float64) string {
	x1, y1 := coords(from)
	x2, y2 := coords(to)

	largeArc := "0"

	if to-from > 0.5 {
		largeArc = "1"
	}

	return fmt.Sprintf("M %.3f %.3f A 1 1 0 %s 1 %.3f %.3f L 0 0", x1, y1, largeArc, x2, y2)
}
