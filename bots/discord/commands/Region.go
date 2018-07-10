package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var regions = map[string]string{
	"africa":    "465876853236826112",
	"america":   "465876808311635979",
	"asia":      "465876834031108096",
	"australia": "465876893036707840",
	"europe":    "465876773029019659",
}

// Region sets the specific region role for the user.
func Region(s *discordgo.Session, msg *discordgo.MessageCreate) bool {
	if strings.HasPrefix(msg.Content, "!region ") {
		return false
	}

	region := strings.ToLower(msg.Content[len("!region "):])

	// Check to make sure the region is in the region map
	if _, ok := regions[region]; !ok {
		s.ChannelMessageSend(msg.ChannelID, "This is not a region!")
		return true
	}

	// Get the channel, this is used to get the guild ID
	c, _ := s.Channel(msg.ChannelID)

	// Check to see if user already has a region role
	user, _ := s.GuildMember(c.GuildID, msg.Author.ID)

	for _, role := range user.Roles {
		match := false

		// We also need to loop through our map because Discord doesn't
		// return roles as names but rather IDs.
		for _, id := range regions {
			if role == id {
				// Remove the role and set match to true
				s.GuildMemberRoleRemove(c.GuildID, msg.Author.ID, id)
				match = true
				break
			}
		}

		if match {
			break
		}
	}

	// Try to set the role
	err := s.GuildMemberRoleAdd(c.GuildID, msg.Author.ID, regions[region])

	if err != nil {
		s.ChannelMessageSend(msg.ChannelID, "The region role could not be set!")
		return true
	}

	s.ChannelMessageSend(msg.ChannelID, "Set region "+region+" for your account!")
	return true
}
