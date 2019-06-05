package commands

import (
	"strings"

	"github.com/akyoto/color"
	"github.com/bwmarrin/discordgo"
)

// AnimeList shows the link for the anime list of a user.
func AnimeList(s *discordgo.Session, msg *discordgo.MessageCreate) bool {
	if !strings.HasPrefix(msg.Content, "!animelist ") {
		return false
	}

	_, err := s.ChannelMessageSend(msg.ChannelID, "https://notify.moe/+"+strings.Split(msg.Content, " ")[1]+"/animelist/watching")

	if err != nil {
		color.Red(err.Error())
	}

	return true
}
