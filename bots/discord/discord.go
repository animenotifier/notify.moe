package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New()

	if err != nil {
		panic(err)
	}

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
	defer func() {
		_, err := discord.ChannelMessageSend(logChannel, "I'm feeling like shit today so I'm shutting down. B-baka!")

		if err != nil {
			color.Red(err.Error())
		}
	}()

	// Receive events
	discord.AddHandler(OnMessageCreate)
	discord.AddHandler(OnGuildMemberAdd)
	discord.AddHandler(OnGuildMemberRemove)

	// Wait for a CTRL-C
	_, err = discord.ChannelMessageSend(logChannel, "Hooray, I'm up again! Did you miss me?")

	if err != nil {
		color.Red(err.Error())
	}

	log.Printf("Tsundere is ready. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
