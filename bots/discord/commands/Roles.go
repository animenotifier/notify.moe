package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Guild ID
const guildID = "134910939140063232"

// Roles prints out all roles.
func Roles(s *discordgo.Session, msg *discordgo.MessageCreate) bool {
	if msg.Content != "!roles" {
		return false
	}

	roles, _ := s.GuildRoles(guildID)

	for _, role := range roles {
		s.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("%s: %s", role.ID, role.Name))
	}

	return true
}
