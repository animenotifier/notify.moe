package main

import "github.com/animenotifier/arn"

var items = []*arn.Item{
	&arn.Item{
		ID:    "pro-account-3",
		Name:  "PRO Account (1 season)",
		Price: 900,
		Description: `PRO account for 1 anime season (3 months).

1 month equals 300 gems.

Includes:

* Special highlight on the forums
* High priority for your personal suggestions
* Access to the VIP channel on Discord
* Early access to new features`,
		Icon:       "star",
		Rarity:     arn.ItemRaritySuperior,
		Order:      1,
		Consumable: true,
	},
	&arn.Item{
		ID:    "pro-account-6",
		Name:  "PRO Account (2 seasons)",
		Price: 1600,
		Description: `PRO account for 2 anime seasons (6 months).

11% less monthly costs compared to standard.

Includes:

* Special highlight on the forums
* High priority for your personal suggestions
* Access to the VIP channel on Discord
* Early access to new features`,
		Icon:       "star",
		Rarity:     arn.ItemRarityRare,
		Order:      2,
		Consumable: true,
	},
	&arn.Item{
		ID:    "pro-account-12",
		Name:  "PRO Account (4 seasons)",
		Price: 3000,
		Description: `PRO account for 4 anime seasons (12 months).

16% less monthly costs compared to standard.

Includes:

* Special highlight on the forums
* High priority for your personal suggestions
* Access to the VIP channel on Discord
* Early access to new features`,
		Icon:       "star",
		Rarity:     arn.ItemRarityUnique,
		Order:      3,
		Consumable: true,
	},
	&arn.Item{
		ID:    "pro-account-24",
		Name:  "PRO Account (8 seasons)",
		Price: 5900,
		Description: `PRO account for 8 anime seasons (24 months).

18% less monthly costs compared to standard.

Includes:

* Special highlight on the forums
* High priority for your personal suggestions
* Access to the VIP channel on Discord
* Early access to new features`,
		Icon:       "star",
		Rarity:     arn.ItemRarityLegendary,
		Order:      4,
		Consumable: true,
	},
	&arn.Item{
		ID:    "anime-support-ticket",
		Name:  "Anime Support Ticket",
		Price: 100,
		Description: `Support the makers of your favourite anime by using an anime support ticket.
Anime Notifier uses 10% of the money to handle the transaction fees while the remaining 90% go directly
to the studios involved in the creation of your favourite anime.

*This feature is work in progress.*`,
		Icon:       "ticket",
		Rarity:     arn.ItemRarityRare,
		Order:      5,
		Consumable: false,
	},
}

func main() {
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
