package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Guild ID
const guildID = "134910939140063232"

// Admin ID
const adminID = "122970452632141826"

// Roles prints out all roles for the server admin.
func Roles(s *discordgo.Session, msg *discordgo.MessageCreate) bool {
	if msg.Content != "!roles" || msg.Author.ID != adminID {
		return false
	}

	roles, _ := s.GuildRoles(guildID)

	for _, role := range roles {
		s.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("%s: %s", role.ID, role.Name))
	}

	return true
}
