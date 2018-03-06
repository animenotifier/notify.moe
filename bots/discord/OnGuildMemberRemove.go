package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// OnGuildMemberRemove is called every time a user leaves the server.
func OnGuildMemberRemove(session *discordgo.Session, event *discordgo.GuildMemberRemove) {
	fmt.Println(event.Member.User.Username + " just left!")

	// session.ChannelMessageSend(welcomeChannel, "Good bye "+event.Member.User.Mention()+"!")
}
