package ability

import (
	"fmt"
	"github.com/mtslzr/pokeapi-go/structs"
)

type Ability struct {
	Name string
	Effect string
}

func MapAbility(orginal structs.Ability) (Ability) {
	ability := Ability{}
	ability.Name = orginal.Name
	ability.Effect = getFirstEnglishEffectEntry(orginal)
	return ability
}

func getFirstEnglishEffectEntry(ability structs.Ability) string {
	for _, entry := range ability.EffectEntries {
		if entry.Language.Name == "en" {
			fmt.Println("found")
			return entry.ShortEffect
		}
	}
	fmt.Println("doh")
	return ""
}