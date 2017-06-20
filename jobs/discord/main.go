package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/animenotifier/arn"
	"github.com/bwmarrin/discordgo"
)

// Session provides access to the Discord session.
var discord *discordgo.Session

func main() {
	var err error

	exe, err := os.Executable()

	if err != nil {
		panic(err)
	}

	dir := path.Dir(exe)
	var apiKeysPath string
	apiKeysPath, err = filepath.Abs(dir + "/../../security/api-keys.json")

	if err != nil {
		panic(err)
	}

	var apiKeys arn.APIKeys
	data, _ := ioutil.ReadFile(apiKeysPath)
	err = json.Unmarshal(data, &apiKeys)

	if err != nil {
		panic(err)
	}

	discord, _ = discordgo.New()
	discord.Token = "Bot " + apiKeys.Discord.Token

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
		userResults, animeResults := arn.Search(term, 10, 10)
		message := ""

		for _, user := range userResults {
			message += "https://notify.moe/" + user.Link() + "\n"
		}

		for _, anime := range animeResults {
			message += "https://notify.moe/" + anime.Link() + "\n"
		}

		if len(userResults) == 0 && len(animeResults) == 0 {
			message = "Sorry, I couldn't find any anime or users with that term."
		}

		s.ChannelMessageSend(m.ChannelID, message)
		return
	}
}
