package pokemon

import (
	"github.com/mtslzr/pokeapi-go"
)

func GetPokemon(name string) (Pokemon) {
	pokemon := getMappedPokemon(name)
	return pokemon
}

func getMappedPokemon(name string) (Pokemon) {	
	pokemonResponse, err := pokeapi.Pokemon(name)
	if err != nil {
		panic("Error getting pokemon")
	}
	return MapPokemon(pokemonResponse)
}