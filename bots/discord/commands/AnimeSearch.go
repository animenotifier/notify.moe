package commands

import (
	"strings"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn/search"
	"github.com/bwmarrin/discordgo"
)

// AnimeSearch shows the link for the anime list of a user.
func AnimeSearch(s *discordgo.Session, msg *discordgo.MessageCreate) bool {
	if !strings.HasPrefix(msg.Content, "!a ") {
		return false
	}

	term := msg.Content[len("!a "):]
	animes := search.Anime(term, 3)
	message := ""

	for _, anime := range animes {
		message += "https://notify.moe" + anime.Link() + "\n"
	}

	if len(animes) == 0 {
		message = "Sorry, I couldn't find anything using that term."
	}

	_, err := s.ChannelMessageSend(msg.ChannelID, message)

	if err != nil {
		color.Red(err.Error())
	}

	return true
}
