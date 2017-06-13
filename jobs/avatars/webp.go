package main

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"strings"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
	"github.com/nfnt/resize"
)

func makeWebPAvatar(user *arn.User) {
	baseName := findOriginalAvatar(user)

	if baseName == "" {
		return
	}

	original := avatarsDirectoryOriginal + baseName
	outFile := avatarsDirectoryWebP + user.ID + ".webp"

	err := avatarToWebP(original, outFile, 80)

	if err != nil {
		color.Red(user.Nick + " [WebP]")
	} else {
		color.Green(user.Nick + " [WebP]")
	}
}

func avatarToWebP(in string, out string, quality float32) error {
	img, _, loadErr := arn.LoadImage(in)

	if loadErr != nil {
		return loadErr
	}

	// Big avatar
	saveErr := arn.SaveWebP(img, out, quality)

	if saveErr != nil {
		return saveErr
	}

	// Small avatar
	smallImg := resize.Resize(arn.AvatarSmallSize, 0, img, resize.Lanczos3)
	saveErr = arn.SaveWebP(smallImg, strings.Replace(out, "large/", "small/", 1), quality)
	return saveErr
}
