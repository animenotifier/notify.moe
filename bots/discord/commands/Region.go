package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var regions = map[string]string{
	"africa":    "465387147629953034",
	"america":   "465386843706359840",
	"asia":      "465386826006528001",
	"australia": "465387169888862230",
	"europe":    "465386794914152448",
}

// Region sets the specific region role for the user.
func Region(s *discordgo.Session, msg *discordgo.MessageCreate) bool {
	if !strings.HasPrefix(msg.Content, "!region ") {
		return false
	}

	region := strings.ToLower(msg.Content[len("!region "):])

	// Check to make sure the region is in the region map
	if _, ok := regions[region]; !ok {
		s.ChannelMessageSend(msg.ChannelID, "This is not a region!")
		return true
	}

	// Check to see if user already has a region role
	user, err := s.GuildMember(guildID, msg.Author.ID)

	if err != nil {
		s.ChannelMessageSend(msg.ChannelID, "Error: User not found")
		return true
	}

	for _, role := range user.Roles {
		match := false

		// We also need to loop through our map because Discord doesn't
		// return roles as names but rather IDs.
		for _, id := range regions {
			if role == id {
				// Remove the role and set match to true
				s.GuildMemberRoleRemove(guildID, msg.Author.ID, id)
				match = true
				break
			}
		}

		if match {
			break
		}
	}

	// Try to set the role
	err = s.GuildMemberRoleAdd(guildID, msg.Author.ID, regions[region])

	if err != nil {
		s.ChannelMessageSend(msg.ChannelID, "The region role could not be set!")
		return true
	}

	s.ChannelMessageSend(msg.ChannelID, "Set region "+region+" for your account!")
	return true
}
