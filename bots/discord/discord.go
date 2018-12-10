package main

import (
	"log"
	"os"
	"os/signal"
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
	defer arn.Node.Close()

	// Receive events
	discord.AddHandler(OnMessageCreate)
	discord.AddHandler(OnGuildMemberAdd)
	discord.AddHandler(OnGuildMemberRemove)

	// Wait for a CTRL-C
	log.Printf("Tsundere is ready. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
