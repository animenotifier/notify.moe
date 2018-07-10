package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// AnimeList shows the link for the anime list of a user.
func AnimeList(s *discordgo.Session, msg *discordgo.MessageCreate) bool {
	if !strings.HasPrefix(msg.Content, "!animelist ") {
		return false
	}

	s.ChannelMessageSend(msg.ChannelID, "https://notify.moe/+"+strings.Split(msg.Content, " ")[1]+"/animelist/watching")
	return true
}
