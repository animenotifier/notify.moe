package explorecolor

import (
	"math"
	"strconv"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils/infinitescroll"
)

const (
	animeFirstLoad = 50
	animePerScroll = 30
)

// AnimeByAverageColor returns all anime with an image in the given color.
func AnimeByAverageColor(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	color := ctx.Get("color")
	index, _ := ctx.GetInt("index")

	allAnimes := filterAnimeByColor(color)
	arn.SortAnimeByQuality(allAnimes)

	// Slice the part that we need
	animes := allAnimes[index:]
	maxLength := animeFirstLoad

	if index > 0 {
		maxLength = animePerScroll
	}

	if len(animes) > maxLength {
		animes = animes[:maxLength]
	}

	// Next index
	nextIndex := infinitescroll.NextIndex(ctx, len(allAnimes), maxLength, index)

	// In case we're scrolling, send animes only (without the page frame)
	if index > 0 {
		return ctx.HTML(components.AnimeGridScrollable(animes, user))
	}

	// Otherwise, send the full page
	return ctx.HTML(components.ExploreColor(animes, nextIndex, len(allAnimes), color, user))
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
