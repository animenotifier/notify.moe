package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// LowResolutionAnimeImages filters anime with low resolution images.
func LowResolutionAnimeImages(ctx aero.Context) error {
	return filterAnimeImages(ctx, "Anime with low resolution images", arn.AnimeImageLargeWidth, arn.AnimeImageLargeHeight)
}

// UltraLowResolutionAnimeImages filters anime with ultra low resolution images.
func UltraLowResolutionAnimeImages(ctx aero.Context) error {
	return filterAnimeImages(ctx, "Anime with ultra low resolution images", arn.AnimeImageLargeWidth/2, arn.AnimeImageLargeHeight/2)
}

func filterAnimeImages(ctx aero.Context, title string, minExpectedWidth int, minExpectedHeight int) error {
	return editorList(
		ctx,
		title,
		func(anime *arn.Anime) bool {
			return anime.Image.Width < minExpectedWidth || anime.Image.Height < minExpectedHeight
		},
		googleImageSearch,
	)
}

func googleImageSearch(anime *arn.Anime) string {
	return "https://www.google.com/search?q=" + anime.Title.Canonical + " anime cover" + "&tbm=isch&tbs=imgo:1,isz:lt,islt:qsvga"
}
