package item

import (
	"github.com/mtslzr/pokeapi-go/structs"
)

type Item struct {
	Name string
	ShortEffect string
	Sprite string
}

func MapItem(original structs.Item) Item {
	item := Item{}
	item.Name = original.Name
	item.ShortEffect = getFirstEnglishEffectEntry(original)
	// item.Sprite = orginal.Sprites
	return item
}

func getFirstEnglishEffectEntry(item structs.Item) string {
	for _, entry := range item.EffectEntries {
		if entry.Language.Name == "en" {
			return entry.ShortEffect
		}
	}
	return ""
}



