package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Channels
const welcomeChannel = "420600996268474368"
const enChannel = "134910939140063232"
const jpChannel = "334016470474424322"

// Emoji
const ramFive = ":ramFive:418520880511975426"
const remFive = ":remFive:403671154553782275"

// OnGuildMemberAdd is called every time a user joins the server.
func OnGuildMemberAdd(session *discordgo.Session, event *discordgo.GuildMemberAdd) {
	fmt.Println(event.Member.User.Username + " just joined!")

	msg, err := session.ChannelMessageSend(welcomeChannel, fmt.Sprintf(
		"**Welcome** %s!\n\nTo join this server, you need to verify your notify.moe account. Simply type `!verify` to receive instructions on how to do that.\n\nAfterwards, please introduce yourself in <#%s>.\n\n日本人は <#%s> で自己紹介して下さい！",
		event.Member.User.Mention(),
		enChannel,
		jpChannel,
	))

	if err != nil {
		fmt.Println(err)
		return
	}

	err = session.MessageReactionAdd(welcomeChannel, msg.ID, ramFive)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = session.MessageReactionAdd(welcomeChannel, msg.ID, remFive)

	if err != nil {
		fmt.Println(err)
		return
	}
}
