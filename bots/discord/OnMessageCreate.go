package main

import (
	"math/rand"
	"strings"

	"github.com/animenotifier/arn"

	"github.com/animenotifier/arn/search"
	"github.com/bwmarrin/discordgo"
)

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

	// Anime search
	if strings.HasPrefix(msg.Content, "!a ") {
		term := msg.Content[len("!a "):]
		animes := search.Anime(term, 3)
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

	// Anime list of user
	if strings.HasPrefix(msg.Content, "!animelist ") {
		s.ChannelMessageSend(msg.ChannelID, "https://notify.moe/+"+strings.Split(msg.Content, " ")[1]+"/animelist/watching")
		return
	}

	// Play status
	if strings.HasPrefix(msg.Content, "!play ") {
		s.UpdateStatus(0, msg.Content[len("!play "):])
		return
	}

	// Random quote
	if msg.Content == "!randomquote" {
		allQuotes := arn.FilterQuotes(func(quote *arn.Quote) bool {
			return !quote.IsDraft && quote.IsValid()
		})

		quote := allQuotes[rand.Intn(len(allQuotes))]
		s.ChannelMessageSend(msg.ChannelID, "https://notify.moe"+quote.Link())
		return
	}

	// GitHub source of the bot
	if msg.Content == "!source" {
		s.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+" B-baaaaaaaka! Y..you...you want to...TOUCH MY CODE?!\n\nhttps://github.com/animenotifier/notify.moe/tree/go/bots/discord")
		return
	}

	// Set the specific region role for the user
	if strings.HasPrefix(msg.Content, "!region ") {

		regions := map[string]string{
			"africa":    "465876853236826112",
			"america":   "465876808311635979",
			"asia":      "465876834031108096",
			"australia": "465876893036707840",
			"europe":    "465876773029019659",
		}

		region := msg.Content[len("!region "):]

		// check to make sure the region is in the region map
		if _, ok := regions[region]; ok {
			// try to set the role
			c, _ := s.Channel(msg.ChannelID)
			err := s.GuildMemberRoleAdd(c.GuildID, msg.Author.ID, regions[region])

			if err != nil {
				s.ChannelMessageSend(msg.ChannelID, "The region role could not be set!")
				return
			}

			s.ChannelMessageSend(msg.ChannelID, "The "+region+" role has been set on your user!")

		} else {
			s.ChannelMessageSend(msg.ChannelID, "This is not a region!")
		}

		return
	}
}
