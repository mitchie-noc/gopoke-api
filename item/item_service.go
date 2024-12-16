package item

import (
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetItemsByCategory(category string) (structs.ItemCategory) {
	items, err := pokeapi.ItemCategory(category)
	if err != nil {
		panic("Error getting item by category")
	}
	return items
}

func GetItemByName(name string) (Item) {
	item, err := pokeapi.Item(name)
	if err != nil {
		panic("Error getting item by name")
	}
	return MapItem(item)
}
