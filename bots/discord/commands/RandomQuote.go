package commands

import (
	"math/rand"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
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
	_, err := s.ChannelMessageSend(msg.ChannelID, "https://notify.moe"+quote.Link())

	if err != nil {
		color.Red(err.Error())
	}

	return true
}
