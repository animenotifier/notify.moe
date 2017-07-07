package utils

import (
	"fmt"
	"math"
)

// PieChart ...
type PieChart struct {
	Title  string
	Slices []*PieChartSlice
}

// PieChartSlice ...
type PieChartSlice struct {
	From  float64
	To    float64
	Title string
	Color string
}

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

// NewPieChart ...
func NewPieChart(title string, data map[string]float64) *PieChart {
	return &PieChart{
		Title:  title,
		Slices: ToPieChartSlices(data),
	}
}

// ToPieChartSlices ...
func ToPieChartSlices(data map[string]float64) []*PieChartSlice {
	if len(data) == 0 {
		return nil
	}

	dataSorted := []*AnalyticsItem{}
	sum := 0.0

	for key, value := range data {
		sum += value

		item := &AnalyticsItem{
			Key:   key,
			Value: value,
		}

		if len(dataSorted) == 0 {
			dataSorted = append(dataSorted, item)
			continue
		}

		found := false

		for i := 0; i < len(dataSorted); i++ {
			if value >= dataSorted[i].Value {
				// Append empty element
				dataSorted = append(dataSorted, nil)

				// Move all elements after index "i" 1 position up
				copy(dataSorted[i+1:], dataSorted[i:])

				// Set value for index "i"
				dataSorted[i] = item

				// Set flag
				found = true

				// Leave loop
				break
			}
		}

		if !found {
			dataSorted = append(dataSorted, item)
		}
	}

	slices := []*PieChartSlice{}
	current := 0.0
	hueOffset := 0.0
	hueScaling := 60.0

	for _, item := range dataSorted {
		percentage := float64(item.Value) / sum

		slices = append(slices, &PieChartSlice{
			From:  current,
			To:    current + percentage,
			Title: fmt.Sprintf("%s (%d%%)", item.Key, int(percentage*100+0.5)),
			Color: fmt.Sprintf("hsl(%.2f, 75%%, 50%%)", current*hueScaling+hueOffset),
		})

		current += percentage
	}

	return slices
}
