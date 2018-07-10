package commands

import (
	"math/rand"

	"github.com/animenotifier/arn"
	"github.com/bwmarrin/discordgo"
)

// RandomQuote shows a random quote.
func RandomQuote(s *discordgo.Session, msg *discordgo.MessageCreate) bool {
	if msg.Content != "!randomquote" {
		return false
	}

	allQuotes := arn.FilterQuotes(func(quote *arn.Quote) bool {
		return !quote.IsDraft && quote.IsValid()
	})

	quote := allQuotes[rand.Intn(len(allQuotes))]
	s.ChannelMessageSend(msg.ChannelID, "https://notify.moe"+quote.Link())
	return true
}
