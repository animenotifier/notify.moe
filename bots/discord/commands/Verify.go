package commands

import (
	"fmt"
	"strings"

	"github.com/animenotifier/arn"
	"github.com/akyoto/color"
	"github.com/pariz/gountries"

	"github.com/bwmarrin/discordgo"
)

var query = gountries.New()

var regions = map[string]string{
	"africa":   "465387147629953034",
	"americas": "465386843706359840",
	"asia":     "465386826006528001",
	"oceania":  "465387169888862230",
	"europe":   "465386794914152448",
}

const (
	verifiedRole  = "512044929195704330"
	editorRole    = "141849207404363776"
	staffRole     = "218221363918274560"
	supporterRole = "365719917426638848"
)

// Verify verifies that the given user has an account on notify.moe.
func Verify(s *discordgo.Session, msg *discordgo.MessageCreate) bool {
	discordTag := msg.Author.Username + "#" + msg.Author.Discriminator

	if msg.Content == "!verify" {
		s.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("To verify your account, first add `%s` as your Discord account on https://notify.moe/settings/accounts, then type `!verify` followed by your username on notify.moe, e.g. `!verify MyName`", discordTag))
		return true
	}

	if !strings.HasPrefix(msg.Content, "!verify ") {
		return false
	}

	arnUserName := strings.TrimSpace(msg.Content[len("!verify "):])
	user, err := arn.GetUserByNick(arnUserName)

	if err != nil {
		s.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("User `%s` doesn't seem to exist on notify.moe", arnUserName))
		return true
	}

	if user.Accounts.Discord.Nick == "" {
		s.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("You haven't set up your Discord account `%s` on https://notify.moe/settings/accounts yet", discordTag))
		return true
	}

	if user.Accounts.Discord.Nick != discordTag {
		s.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("Mismatching Discord accounts: `%s` and `%s`", user.Accounts.Discord.Nick, discordTag))
		return true
	}

	// Try to set the verified role
	err = s.GuildMemberRoleAdd(guildID, msg.Author.ID, verifiedRole)

	if err != nil {
		s.ChannelMessageSend(msg.ChannelID, "There was an error adding the Verified role to your account!")
		return true
	}

	// Give editor role
	if user.Role == "editor" {
		s.GuildMemberRoleAdd(guildID, msg.Author.ID, editorRole)
		s.GuildMemberRoleAdd(guildID, msg.Author.ID, staffRole)
	}

	// Give region role
	if user.Location.CountryCode != "" {
		country, err := query.FindCountryByAlpha(user.Location.CountryCode)

		if err != nil {
			color.Red("Error querying country code: %s", err.Error())
		} else {
			regionRole, exists := regions[strings.ToLower(country.Region)]

			if !exists {
				color.Red("Error getting region role for: %s", country.Region)
			} else {
				// Remove old region role
				for _, roleID := range regions {
					s.GuildMemberRoleRemove(guildID, msg.Author.ID, roleID)
				}

				// Add new region role
				s.GuildMemberRoleAdd(guildID, msg.Author.ID, regionRole)
			}
		}
	}

	// Give or remove supporter role
	if user.IsPro() {
		s.GuildMemberRoleAdd(guildID, msg.Author.ID, supporterRole)
	} else {
		s.GuildMemberRoleRemove(guildID, msg.Author.ID, supporterRole)
	}

	// Update nickname to notify.moe nick
	err = s.GuildMemberNickname(guildID, msg.Author.ID, user.Nick)

	if err != nil {
		color.Red("Error updating nickname: %s", err.Error())
	}

	// Update notify.moe user account
	if !user.Accounts.Discord.Verified {
		user.Accounts.Discord.Verified = true
		user.Save()
	}

	// Send success message
	s.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("%s Thank you, you are now a verified member of the notify.moe community!", msg.Author.Mention()))
	return true
}
