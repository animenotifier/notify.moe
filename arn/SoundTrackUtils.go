package arn

import (
	"errors"
	"regexp"
)

var youtubeIDRegex = regexp.MustCompile(`youtu(?:.*\/v\/|.*v=|\.be\/)([A-Za-z0-9_-]{11})`)

// GetYoutubeMedia returns an ExternalMedia object for the given Youtube link.
func GetYoutubeMedia(url string) (*ExternalMedia, error) {
	matches := youtubeIDRegex.FindStringSubmatch(url)

	if len(matches) < 2 {
		return nil, errors.New("Invalid Youtube URL")
	}

	videoID := matches[1]

	media := &ExternalMedia{
		Service:   "Youtube",
		ServiceID: videoID,
	}

	return media, nil
}
