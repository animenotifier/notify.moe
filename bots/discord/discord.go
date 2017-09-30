package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/animenotifier/arn"
	"github.com/bwmarrin/discordgo"
)

// Session provides access to the Discord session.
var discord *discordgo.Session

func main() {
	var err error

	discord, _ = discordgo.New()
	discord.Token = "Bot " + arn.APIKeys.Discord.Token

	// Verify a Token was provided
	if discord.Token == "" {
		log.Println("You must provide a Discord authentication token.")
		return
	}

	// Verify the Token is valid and grab user information
	discord.State.User, err = discord.User("@me")

	if err != nil {
		log.Printf("Error fetching user information: %s\n", err)
	}

	// Open a websocket connection to Discord
	err = discord.Open()

	if err != nil {
		log.Printf("Error opening connection to Discord, %s\n", err)
	}

	defer discord.Close()

	// Receive messages
	discord.AddHandler(onMessage)

	// Wait for a CTRL-C
	log.Printf("Tsundere is ready. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

// This function will be called every time a new message is created on any channel.
func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!commands" {
		s.ChannelMessageSend(m.ChannelID, `
**!user** [username]
**!anime** [id]
**!animelist** [username]
**!tag** [forum tag]`)
	}

	// Has the bot been mentioned?
	for _, user := range m.Mentions {
		if user.ID == discord.State.User.ID {
			s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" :heart:")
			return
		}
	}

	if strings.HasPrefix(m.Content, "!user ") {
		s.ChannelMessageSend(m.ChannelID, "https://notify.moe/+"+strings.Split(m.Content, " ")[1])
		return
	}

	if strings.HasPrefix(m.Content, "!anime ") {
		s.ChannelMessageSend(m.ChannelID, "https://notify.moe/anime/"+strings.Split(m.Content, " ")[1])
		return
	}

	if strings.HasPrefix(m.Content, "!animelist ") {
		s.ChannelMessageSend(m.ChannelID, "https://notify.moe/+"+strings.Split(m.Content, " ")[1]+"/animelist")
		return
	}

	if strings.HasPrefix(m.Content, "!tag ") {
		s.ChannelMessageSend(m.ChannelID, "https://notify.moe/forum/"+strings.ToLower(strings.Split(m.Content, " ")[1]))
		return
	}

	if strings.HasPrefix(m.Content, "!s ") {
		term := m.Content[len("!s "):]
		users, animes, posts, threads := arn.Search(term, 3, 3, 3, 3)
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

		if len(users) == 0 && len(animes) == 0 && len(posts) == 0 && len(threads) == 0 {
			message = "Sorry, I couldn't find anything using that term."
		}

		s.ChannelMessageSend(m.ChannelID, message)
		return
	}
}
