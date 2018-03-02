package main

import (
	"fmt"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/arn/mailer"
)

func main() {
	notification := &arn.PushNotification{
		Title:   "Boku dake ga Inai Machi",
		Message: "Episode 16 has been released!",
		Icon:    "https://media.notify.moe/images/anime/large/11110.webp",
		Link:    "https://notify.moe/anime/11110",
	}

	err := mailer.SendEmailNotification("e.urbach@gmail.com", notification)

	if err != nil {
		fmt.Println(err)
	}
}
