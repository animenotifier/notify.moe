package main

import (
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/bots/discord/commands"

	"github.com/bwmarrin/discordgo"
)

// Command represents a single bot command function signature.
type Command func(*discordgo.Session, *discordgo.MessageCreate) bool

var allCommands = []Command{
	commands.AnimeList,
	commands.AnimeSearch,
	commands.Did,
	commands.Play,
	commands.RandomQuote,
	commands.Roles,
	commands.Source,
	commands.Verify,
}

// OnMessageCreate is called every time a new message is created on any channel.
func OnMessageCreate(s *discordgo.Session, msg *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if msg.Author.ID == s.State.User.ID {
		return
	}

	if msg.Content == "!help" || msg.Content == "!commands" {
		_, err := s.ChannelMessageSend(msg.ChannelID, `
**!a** [anime search term]
**!animelist** [username]
**!play** [status text]
**!randomquote**
**!source**
**!verify** [username]`)

		if err != nil {
			color.Red(err.Error())
		}
	}

	// Has the bot been mentioned?
	// for _, user := range msg.Mentions {
	// 	if user.ID == discord.State.User.ID {
	// 		s.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+" :heart:")
	// 		return
	// 	}
	// }

	// Has the user invoked a command?
	for _, cmd := range allCommands {
		if cmd(s, msg) {
			return
		}
	}
}
