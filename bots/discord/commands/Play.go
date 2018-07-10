package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Play changes the status of the bot.
func Play(s *discordgo.Session, msg *discordgo.MessageCreate) bool {
	if !strings.HasPrefix(msg.Content, "!play ") {
		return false
	}

	s.UpdateStatus(0, msg.Content[len("!play "):])
	return true
}
