package autocorrect

import (
	"regexp"
	"strings"
)

const maxNickLength = 25

var fixNickRegex = regexp.MustCompile(`[\W\s\d]`)

var accountNickRegexes = []*regexp.Regexp{
	regexp.MustCompile(`anilist.co/user/(.*)`),
	regexp.MustCompile(`anilist.co/animelist/(.*)`),
	regexp.MustCompile(`kitsu.io/users/(.*?)/library`),
	regexp.MustCompile(`kitsu.io/users/(.*)`),
	regexp.MustCompile(`anime-planet.com/users/(.*?)/anime`),
	regexp.MustCompile(`anime-planet.com/users/(.*)`),
	regexp.MustCompile(`myanimelist.net/profile/(.*)`),
	regexp.MustCompile(`myanimelist.net/animelist/(.*?)\?`),
	regexp.MustCompile(`myanimelist.net/animelist/(.*)`),
	regexp.MustCompile(`myanimelist.net/(.*)`),
	regexp.MustCompile(`myanimelist.com/(.*)`),
	regexp.MustCompile(`twitter.com/(.*)`),
	regexp.MustCompile(`osu.ppy.sh/u/(.*)`),
}

var animeLinkRegex = regexp.MustCompile(`notify.moe/anime/(\d+)`)
var osuBeatmapRegex = regexp.MustCompile(`osu.ppy.sh/s/(\d+)`)

// Tag converts links to correct tags automatically.
func Tag(tag string) string {
	tag = strings.TrimSpace(tag)
	tag = strings.TrimSuffix(tag, "/")

	// Anime
	matches := animeLinkRegex.FindStringSubmatch(tag)

	if len(matches) > 1 {
		return "anime:" + matches[1]
	}

	// Osu beatmap
	matches = osuBeatmapRegex.FindStringSubmatch(tag)

	if len(matches) > 1 {
		return "osu-beatmap:" + matches[1]
	}

	return tag
}

// UserNick automatically corrects a username.
func UserNick(nick string) string {
	nick = fixNickRegex.ReplaceAllString(nick, "")

	if nick == "" {
		return nick
	}

	nick = strings.Trim(nick, "_")

	if nick == "" {
		return ""
	}

	if len(nick) > maxNickLength {
		nick = nick[:maxNickLength]
	}

	return strings.ToUpper(string(nick[0])) + nick[1:]
}

// AccountNick automatically corrects the username/nick of an account.
func AccountNick(nick string) string {
	for _, regex := range accountNickRegexes {
		matches := regex.FindStringSubmatch(nick)

		if len(matches) > 1 {
			nick = matches[1]
			return nick
		}
	}

	return nick
}

// PostText fixes common mistakes in post texts.
func PostText(text string) string {
	text = strings.Replace(text, "http://", "https://", -1)
	text = strings.TrimSpace(text)
	return text
}

// ThreadTitle ...
func ThreadTitle(title string) string {
	return strings.TrimSpace(title)
}

// Website fixed common website mistakes.
func Website(url string) string {
	// Disallow links that aren't actual websites,
	// just tracker links.
	if IsTrackerLink(url) {
		return ""
	}

	url = strings.TrimSpace(url)
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimSuffix(url, "/")

	return url
}

// IsTrackerLink returns true if the URL is a tracker link.
func IsTrackerLink(url string) bool {
	return strings.Contains(url, "myanimelist.net/") || strings.Contains(url, "anilist.co/") || strings.Contains(url, "kitsu.io/") || strings.Contains(url, "kissanime.")
}
