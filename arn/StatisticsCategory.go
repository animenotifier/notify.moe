package arn

import "fmt"

// StatisticsCategory ...
type StatisticsCategory struct {
	Name      string      `json:"name"`
	PieCharts []*PieChart `json:"pieCharts"`
}

// PieChart ...
type PieChart struct {
	Title  string           `json:"title"`
	Slices []*PieChartSlice `json:"slices"`
}

// PieChartSlice ...
type PieChartSlice struct {
	From  float64 `json:"from"`
	To    float64 `json:"to"`
	Title string  `json:"title"`
	Color string  `json:"color"`
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
	hueOffset := 230.0
	hueScaling := -30.0

	for i, item := range dataSorted {
		percentage := item.Value / sum

		slices = append(slices, &PieChartSlice{
			From:  current,
			To:    current + percentage,
			Title: fmt.Sprintf("%s (%d%%)", item.Key, int(percentage*100+0.5)),
			Color: fmt.Sprintf("hsl(%.2f, 75%%, 50%%)", float64(i)*hueScaling+hueOffset),
		})

		current += percentage
	}

	return slices
}
