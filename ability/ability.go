package ability

import (
	"github.com/mtslzr/pokeapi-go"
)

func GetAbility(name string) (Ability) {
	pokemonResponse, err := pokeapi.Ability(name)
	if err != nil {
		panic("Error getting pokemon")
	}
	return MapAbility(pokemonResponse)
}