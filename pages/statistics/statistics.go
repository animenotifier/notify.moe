package statistics

import (
	"fmt"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get ...
func Get(ctx *aero.Context) string {
	analytics, err := arn.AllAnalytics()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Couldn't retrieve analytics", err)
	}

	screenSizes := map[string]int{}

	for _, info := range analytics {
		size := arn.ToString(info.Screen.Width) + " x " + arn.ToString(info.Screen.Height)
		screenSizes[size]++
	}

	screenSizesSorted := []*utils.AnalyticsItem{}

	for size, count := range screenSizes {
		item := &utils.AnalyticsItem{
			Key:   size,
			Value: count,
		}

		if len(screenSizesSorted) == 0 {
			screenSizesSorted = append(screenSizesSorted, item)
			continue
		}

		found := false

		for i := 0; i < len(screenSizesSorted); i++ {
			if count >= screenSizesSorted[i].Value {
				// Append empty element
				screenSizesSorted = append(screenSizesSorted, nil)

				// Move all elements after index "i" 1 position up
				copy(screenSizesSorted[i+1:], screenSizesSorted[i:])

				// Set value for index "i"
				screenSizesSorted[i] = item

				// Set flag
				found = true

				// Leave loop
				break
			}
		}

		if !found {
			screenSizesSorted = append(screenSizesSorted, item)
		}
	}

	slices := []*utils.PieChartSlice{}
	current := 0.0

	for _, item := range screenSizesSorted {
		percentage := float64(item.Value) / float64(len(analytics))

		slices = append(slices, &utils.PieChartSlice{
			From:  current,
			To:    current + percentage,
			Title: fmt.Sprintf("%s (%d%%)", item.Key, int(percentage*100+0.5)),
			Color: fmt.Sprintf("rgba(255, 64, 0, %.3f)", 0.8-current*0.8),
		})

		current += percentage
	}

	return ctx.HTML(components.Statistics(slices))
}
