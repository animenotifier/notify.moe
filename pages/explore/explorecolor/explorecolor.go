package explorecolor

import (
	"math"
	"strconv"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// AnimeByAverageColor returns all anime with an image in the given color.
func AnimeByAverageColor(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	color := ctx.Get("color")
	animes := filterAnimeByColor(color)

	arn.SortAnimeByQuality(animes)

	return ctx.HTML(components.ExploreColor(animes, color, user))
}

func filterAnimeByColor(colorText string) []*arn.Anime {
	if !strings.HasPrefix(colorText, "hsl:") {
		return nil
	}

	colorText = colorText[len("hsl:"):]
	parts := strings.Split(colorText, ",")

	if len(parts) != 3 {
		return nil
	}

	hue, err := strconv.ParseFloat(parts[0], 64)

	if err != nil {
		return nil
	}

	saturation, err := strconv.ParseFloat(parts[1], 64)

	if err != nil {
		return nil
	}

	lightness, err := strconv.ParseFloat(parts[2], 64)

	if err != nil {
		return nil
	}

	color := arn.HSLColor{
		Hue:        hue,
		Saturation: saturation,
		Lightness:  lightness,
	}

	return arn.FilterAnime(func(anime *arn.Anime) bool {
		animeColor := anime.Image.AverageColor
		hueDifference := color.Hue - animeColor.Hue
		saturationDifference := color.Saturation - animeColor.Saturation
		lightnessDifference := color.Lightness - animeColor.Lightness

		return math.Abs(hueDifference) < 0.05 && math.Abs(saturationDifference) < 0.125 && math.Abs(lightnessDifference) < 0.25
	})
}
