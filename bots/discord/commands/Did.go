package commands

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/arn/search"
	"github.com/bwmarrin/discordgo"
)

var (
	watchedAnimeRegex = regexp.MustCompile(`did (.*?) watch (.*?)\?`)
)

// Did answers some questions with the pattern "Did ...?".
func Did(s *discordgo.Session, msg *discordgo.MessageCreate) bool {
	if !strings.HasPrefix(msg.Content, "!did ") {
		return false
	}

	matches := watchedAnimeRegex.FindStringSubmatch(msg.Content)

	if len(matches) < 3 {
		_, err := s.ChannelMessageSend(msg.ChannelID, "I don't understand that question")

		if err != nil {
			color.Red(err.Error())
		}

		return true
	}

	userName := matches[1]
	animeName := matches[2]

	user, err := arn.GetUserByNick(userName)

	if err != nil {
		_, err := s.ChannelMessageSend(msg.ChannelID, "User not found")

		if err != nil {
			color.Red(err.Error())
		}

		return true
	}

	results := search.Anime(animeName, 1)

	if len(results) == 0 {
		_, err := s.ChannelMessageSend(msg.ChannelID, "Anime not found")

		if err != nil {
			color.Red(err.Error())
		}

		return true
	}

	anime := results[0]
	animeList := user.AnimeList()
	listItem := animeList.Find(anime.ID)

	if listItem != nil && listItem.Status == arn.AnimeListStatusCompleted {
		_, err := s.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("Yes, %s has watched %s.", user.Nick, anime.Title.Canonical))

		if err != nil {
			color.Red(err.Error())
		}
	} else {
		_, err := s.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("No, %s hasn't watched %s.", user.Nick, anime.Title.Canonical))

		if err != nil {
			color.Red(err.Error())
		}
	}

	return true
}
