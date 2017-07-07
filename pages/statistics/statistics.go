package statistics

import (
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

	if len(screenSizesSorted) > 5 {
		screenSizesSorted = screenSizesSorted[:5]
	}

	return ctx.HTML(components.Statistics(screenSizesSorted))
}
