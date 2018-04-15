package main

import (
	"strings"

	"github.com/animenotifier/arn"

	"github.com/bwmarrin/discordgo"
)

// OnMessageCreate is called every time a new message is created on any channel.
func OnMessageCreate(s *discordgo.Session, msg *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if msg.Author.ID == s.State.User.ID {
		return
	}

	if msg.Content == "!commands" {
		s.ChannelMessageSend(msg.ChannelID, `
**!user** [username]
**!anime** [id]
**!animelist** [username]
**!tag** [forum tag]`)
	}

	// Has the bot been mentioned?
	for _, user := range msg.Mentions {
		if user.ID == discord.State.User.ID {
			s.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+" :heart:")
			return
		}
	}

	if strings.HasPrefix(msg.Content, "!user ") {
		s.ChannelMessageSend(msg.ChannelID, "https://notify.moe/+"+strings.Split(msg.Content, " ")[1])
		return
	}

	if strings.HasPrefix(msg.Content, "!animelist ") {
		s.ChannelMessageSend(msg.ChannelID, "https://notify.moe/+"+strings.Split(msg.Content, " ")[1]+"/animelist")
		return
	}

	if strings.HasPrefix(msg.Content, "!tag ") {
		s.ChannelMessageSend(msg.ChannelID, "https://notify.moe/forum/"+strings.ToLower(strings.Split(msg.Content, " ")[1]))
		return
	}

	if strings.HasPrefix(msg.Content, "!play ") {
		s.UpdateStatus(0, msg.Content[len("!play "):])
		return
	}

	if msg.Content == "!source" {
		s.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+" B-baaaaaaaka! Y..you...you want to...TOUCH MY CODE?!\n\nhttps://github.com/animenotifier/notify.moe/tree/go/bots/discord")
		return
	}

	if strings.HasPrefix(msg.Content, "!a ") {
		term := msg.Content[len("!a "):]
		animes := arn.SearchAnime(term, 3)
		message := ""

		for _, anime := range animes {
			message += "https://notify.moe" + anime.Link() + "\n"
		}

		if len(animes) == 0 {
			message = "Sorry, I couldn't find anything using that term."
		}

		s.ChannelMessageSend(msg.ChannelID, message)
		return
	}
}
