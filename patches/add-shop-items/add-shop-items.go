package main

import "github.com/animenotifier/notify.moe/arn"

var items = []*arn.ShopItem{
	// 1 month
	{
		ID:    "pro-account-1",
		Name:  "PRO Account - 1 month",
		Price: 300,
		Description: `PRO status for 1 month.

1 month equals 300 gems.

Includes:

* Dark theme for the website and extension
* Upload your own cover image
* PRO star on your profile
* Special highlight on the forums
* Access to the VIP channel on Discord
* Early access to new features`,
		Icon:       "star",
		Rarity:     arn.ShopItemRaritySuperior,
		Order:      1,
		Consumable: true,
	},

	// 3 months
	{
		ID:    "pro-account-3",
		Name:  "PRO Account - 3 months",
		Price: 900,
		Description: `PRO status for 1 anime season (3 months).

1 month equals 300 gems.

Includes:

* Dark theme for the website and extension
* Upload your own cover image
* PRO star on your profile
* Special highlight on the forums
* Access to the VIP channel on Discord
* Early access to new features`,
		Icon:       "star",
		Rarity:     arn.ShopItemRaritySuperior,
		Order:      2,
		Consumable: true,
	},

	// 6 months
	{
		ID:    "pro-account-6",
		Name:  "PRO Account - 6 months",
		Price: 1600,
		Description: `PRO status for 2 anime seasons (6 months).

11% less monthly costs compared to 1 season.

Includes:

* Dark theme for the website and extension
* Upload your own cover image
* PRO star on your profile
* Special highlight on the forums
* Access to the VIP channel on Discord
* Early access to new features`,
		Icon:       "star",
		Rarity:     arn.ShopItemRarityRare,
		Order:      3,
		Consumable: true,
	},
	{
		ID:    "pro-account-12",
		Name:  "PRO Account - 1 year",
		Price: 3000,
		Description: `PRO status for 4 anime seasons (12 months).

16% less monthly costs compared to 1 season.

Includes:

* Dark theme for the website and extension
* Upload your own cover image
* PRO star on your profile
* Special highlight on the forums
* Access to the VIP channel on Discord
* Early access to new features`,
		Icon:       "star",
		Rarity:     arn.ShopItemRarityUnique,
		Order:      4,
		Consumable: true,
	},
	{
		ID:    "pro-account-24",
		Name:  "PRO Account - 2 years",
		Price: 5900,
		Description: `PRO status for 8 anime seasons (24 months).

18% less monthly costs compared to 1 season.

Includes:

* Dark theme for the website and extension
* Upload your own cover image
* PRO star on your profile
* Special highlight on the forums
* Access to the VIP channel on Discord
* Early access to new features`,
		Icon:       "star",
		Rarity:     arn.ShopItemRarityLegendary,
		Order:      5,
		Consumable: true,
	},
	// 	&arn.ShopItem{
	// 		ID:    "anime-support-ticket",
	// 		Name:  "Anime Support Ticket",
	// 		Price: 100,
	// 		Description: `Support the makers of your favourite anime by using an anime support ticket.
	// Anime Notifier uses 15% of the money to handle the transaction fees while the remaining 85% go directly
	// to the studios involved in the creation of your favourite anime.

	// *This feature is work in progress.*`,
	// 		Icon:       "ticket",
	// 		Rarity:     arn.ShopItemRarityRare,
	// 		Order:      5,
	// 		Consumable: false,
	// 	},
}

func main() {
	defer arn.Node.Close()

	for _, item := range items {
		item.Save()
	}
}

//- ShopItem("PRO Account", "6 months", "1600", "star", strings.Replace(strings.Replace(proAccountMarkdown, "3 months", "6 months", 1), "1 anime season", "2 anime seasons", 1))
//- ShopItem("PRO Account", "1 year", "3000", "star", strings.Replace(strings.Replace(proAccountMarkdown, "3 months", "12 months", 1), "1 anime season", "4 anime seasons", 1))
//- ShopItem("PRO Account", "2 years", "5900", "star", strings.Replace(strings.Replace(proAccountMarkdown, "3 months", "24 months", 1), "1 anime season", "8 anime seasons", 1))
//- ShopItem("Anime Support Ticket", "", "100", "ticket", "Support the makers of your favourite anime by using an anime support ticket. Anime Notifier uses 8% of the money to handle the transaction fees while the remaining 92% go directly to the studios involved in the creation of your favourite anime.")
//- ShopItem("Artwork Support Ticket", "", "100", "ticket", "Support the makers of your favourite artwork by using an artwork support ticket. Anime Notifier uses 8% of the money to handle the transaction fees while the remaining 92% go directly to the creator.")
//- ShopItem("Soundtrack Support Ticket", "", "100", "ticket", "Support the makers of your favourite soundtrack by using a soundtrack support ticket. Anime Notifier uses 8% of the money to handle the transaction fees while the remaining 92% go directly to the creator.")
//- ShopItem("AMV Support Ticket", "", "100", "ticket", "Support the makers of your favourite AMV by using an AMV support ticket. Anime Notifier uses 8% of the money to handle the transaction fees while the remaining 92% go directly to the creator.")
