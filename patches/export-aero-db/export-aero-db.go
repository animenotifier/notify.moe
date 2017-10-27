package main

import (
	"time"

	"github.com/aerogo/database"
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	arn.DB.SetScanPriority("high")

	aeroDB := database.New("arn", arn.DBTypes)
	defer aeroDB.Close()

	for typeName := range arn.DB.Types() {
		count := 0

		switch typeName {
		case "Anime":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.Anime) {
				aeroDB.Set(typeName, obj.ID, obj)
				count++
			}

		case "AnimeEpisodes":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.AnimeEpisodes) {
				aeroDB.Set(typeName, obj.AnimeID, obj)
				count++
			}

		case "AnimeList":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.AnimeList) {
				aeroDB.Set(typeName, obj.UserID, obj)
				count++
			}

		case "AnimeCharacters":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.AnimeCharacters) {
				aeroDB.Set(typeName, obj.AnimeID, obj)
				count++
			}

		case "AnimeRelations":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.AnimeRelations) {
				aeroDB.Set(typeName, obj.AnimeID, obj)
				count++
			}

		case "Character":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.Character) {
				aeroDB.Set(typeName, obj.ID, obj)
				count++
			}

		case "Purchase":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.Purchase) {
				aeroDB.Set(typeName, obj.ID, obj)
				count++
			}

		case "PushSubscriptions":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.PushSubscriptions) {
				aeroDB.Set(typeName, obj.UserID, obj)
				count++
			}

		case "User":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.User) {
				aeroDB.Set(typeName, obj.ID, obj)
				count++
			}

		case "Post":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.Post) {
				aeroDB.Set(typeName, obj.ID, obj)
				count++
			}

		case "Thread":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.Thread) {
				aeroDB.Set(typeName, obj.ID, obj)
				count++
			}

		case "Analytics":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.Analytics) {
				aeroDB.Set(typeName, obj.UserID, obj)
				count++
			}

		case "SoundTrack":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.SoundTrack) {
				aeroDB.Set(typeName, obj.ID, obj)
				count++
			}

		case "Item":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.Item) {
				aeroDB.Set(typeName, obj.ID, obj)
				count++
			}

		case "Inventory":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.Inventory) {
				aeroDB.Set(typeName, obj.UserID, obj)
				count++
			}

		case "Settings":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.Settings) {
				aeroDB.Set(typeName, obj.UserID, obj)
				count++
			}

		case "UserFollows":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.UserFollows) {
				aeroDB.Set(typeName, obj.UserID, obj)
				count++
			}

		case "PayPalPayment":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.PayPalPayment) {
				aeroDB.Set(typeName, obj.ID, obj)
				count++
			}

		case "AniListToAnime":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.AniListToAnime) {
				aeroDB.Set(typeName, obj.ServiceID, obj)
				count++
			}

		case "MyAnimeListToAnime":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.MyAnimeListToAnime) {
				aeroDB.Set(typeName, obj.ServiceID, obj)
				count++
			}

		case "SearchIndex":
			anime, _ := arn.DB.Get(typeName, "Anime")
			aeroDB.Set(typeName, "Anime", anime)

			users, _ := arn.DB.Get(typeName, "User")
			aeroDB.Set(typeName, "User", users)

			posts, _ := arn.DB.Get(typeName, "Post")
			aeroDB.Set(typeName, "Post", posts)

			threads, _ := arn.DB.Get(typeName, "Thread")
			aeroDB.Set(typeName, "Thread", threads)

			count += 4

		case "DraftIndex":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.DraftIndex) {
				aeroDB.Set(typeName, obj.UserID, obj)
				count++
			}

		case "EmailToUser":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.EmailToUser) {
				if obj.Email == "" {
					continue
				}

				aeroDB.Set(typeName, obj.Email, obj)
				count++
			}

		case "FacebookToUser":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.FacebookToUser) {
				if obj.ID == "" {
					continue
				}

				aeroDB.Set(typeName, obj.ID, obj)
				count++
			}

		case "GoogleToUser":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.GoogleToUser) {
				if obj.ID == "" {
					continue
				}

				aeroDB.Set(typeName, obj.ID, obj)
				count++
			}

		case "TwitterToUser":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.TwitterToUser) {
				if obj.ID == "" {
					continue
				}

				aeroDB.Set(typeName, obj.ID, obj)
				count++
			}

		case "NickToUser":
			channel, _ := arn.DB.All(typeName)

			for obj := range channel.(chan *arn.NickToUser) {
				if obj.Nick == "" {
					continue
				}

				aeroDB.Set(typeName, obj.Nick, obj)
				count++
			}

		default:
			color.Yellow("Skipping %s", typeName)
			continue
		}

		color.Green("Export %d %s", count, typeName)
	}

	time.Sleep(1 * time.Second)
}
