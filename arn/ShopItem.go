package arn

import "github.com/aerogo/nano"

const (
	// ShopItemRarityCommon ...
	ShopItemRarityCommon = "common"

	// ShopItemRaritySuperior ...
	ShopItemRaritySuperior = "superior"

	// ShopItemRarityRare ...
	ShopItemRarityRare = "rare"

	// ShopItemRarityUnique ...
	ShopItemRarityUnique = "unique"

	// ShopItemRarityLegendary ...
	ShopItemRarityLegendary = "legendary"
)

// ShopItem is a purchasable item in the shop.
type ShopItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Icon        string `json:"icon"`
	Rarity      string `json:"rarity"`
	Order       int    `json:"order"`
	Consumable  bool   `json:"consumable"`
}

// GetShopItem ...
func GetShopItem(id string) (*ShopItem, error) {
	obj, err := DB.Get("ShopItem", id)

	if err != nil {
		return nil, err
	}

	return obj.(*ShopItem), nil
}

// StreamShopItems returns a stream of all shop items.
func StreamShopItems() <-chan *ShopItem {
	channel := make(chan *ShopItem, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("ShopItem") {
			channel <- obj.(*ShopItem)
		}

		close(channel)
	}()

	return channel
}

// AllShopItems returns a slice of all items.
func AllShopItems() ([]*ShopItem, error) {
	all := make([]*ShopItem, 0, DB.Collection("ShopItem").Count())

	for obj := range StreamShopItems() {
		all = append(all, obj)
	}

	return all, nil
}
