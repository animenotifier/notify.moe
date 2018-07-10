package main

import (
	"github.com/animenotifier/notify.moe/bots/discord/commands"

	"github.com/bwmarrin/discordgo"
)

// Command represents a single bot command function signature.
type Command func(*discordgo.Session, *discordgo.MessageCreate) bool

var allCommands = []Command{
	commands.AnimeList,
	commands.AnimeSearch,
	commands.Play,
	commands.RandomQuote,
	commands.Region,
	commands.Roles,
	commands.Source,
}

// OnMessageCreate is called every time a new message is created on any channel.
func OnMessageCreate(s *discordgo.Session, msg *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if msg.Author.ID == s.State.User.ID {
		return
	}

	if msg.Content == "!help" || msg.Content == "!commands" {
		s.ChannelMessageSend(msg.ChannelID, `
**!a** [anime search term]
**!animelist** [username]
**!play** [status text]
**!randomquote**
**!source**
**!region** [region]`)
	}

	// Has the bot been mentioned?
	for _, user := range msg.Mentions {
		if user.ID == discord.State.User.ID {
			s.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+" :heart:")
			return
		}
	}

	// Has the user invoked a command?
	for _, cmd := range allCommands {
		if cmd(s, msg) {
			return
		}
	}

	return
}
