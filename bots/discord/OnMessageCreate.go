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

	if strings.HasPrefix(msg.Content, "!anime ") {
		s.ChannelMessageSend(msg.ChannelID, "https://notify.moe/anime/"+strings.Split(msg.Content, " ")[1])
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

	if strings.HasPrefix(msg.Content, "!s ") {
		term := msg.Content[len("!s "):]
		users, animes, posts, threads, tracks, characters := arn.Search(term, 3, 3, 3, 3, 3, 3)
		message := ""

		for _, user := range users {
			message += "https://notify.moe" + user.Link() + "\n"
		}

		for _, anime := range animes {
			message += "https://notify.moe" + anime.Link() + "\n"
		}

		for _, post := range posts {
			message += "https://notify.moe" + post.Link() + "\n"
		}

		for _, thread := range threads {
			message += "https://notify.moe" + thread.Link() + "\n"
		}

		for _, track := range tracks {
			message += "https://notify.moe" + track.Link() + "\n"
		}

		for _, character := range characters {
			message += "https://notify.moe" + character.Link() + "\n"
		}

		if len(users) == 0 && len(animes) == 0 && len(posts) == 0 && len(threads) == 0 && len(tracks) == 0 {
			message = "Sorry, I couldn't find anything using that term."
		}

		s.ChannelMessageSend(msg.ChannelID, message)
		return
	}
}
