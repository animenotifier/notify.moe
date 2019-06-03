package arn

import (
	"fmt"
	"regexp"
	"strings"
)

type nyaaAnimeProvider struct{}

// Nyaa anime provider (singleton)
var Nyaa = new(nyaaAnimeProvider)

var nyaaInvalidCharsRegex = regexp.MustCompile(`[^[:alnum:]!']`)
var nyaaTVRegex = regexp.MustCompile(` \(?TV\)?`)

// GetLink retrieves the Nyaa title for the given anime
func (nyaa *nyaaAnimeProvider) GetLink(anime *Anime, additionalSearchTerm string) string {
	searchTitle := nyaa.GetTitle(anime) + "+" + additionalSearchTerm
	searchTitle = strings.Replace(searchTitle, " ", "+", -1)

	quality := ""
	subs := ""

	nyaaSuffix := fmt.Sprintf("?f=0&c=1_2&q=%s+%s+%s&s=seeders&o=desc", searchTitle, quality, subs)
	nyaaSuffix = strings.Replace(nyaaSuffix, "++", "+", -1)

	return "https://nyaa.si/" + nyaaSuffix
}

// GetTitle retrieves the Nyaa title for the given anime
func (nyaa *nyaaAnimeProvider) GetTitle(anime *Anime) string {
	return nyaa.BuildTitle(anime.Title.Canonical)
}

// BuildTitle tries to create a title for use on Nyaa
func (nyaa *nyaaAnimeProvider) BuildTitle(title string) string {
	if title == "" {
		return ""
	}

	title = nyaaInvalidCharsRegex.ReplaceAllString(title, " ")
	title = nyaaTVRegex.ReplaceAllString(title, "")
	title = strings.Replace(title, "  ", " ", -1)
	title = strings.TrimSpace(title)

	return title
}
