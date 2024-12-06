package service

import (
	"github.com/mtslzr/pokeapi-go"
	"ryantest.com/learningGo/mapper"
)

func GetPokemon(name string) (mapper.Pokemon) {	
	pokemonResponse, err := pokeapi.Pokemon(name)
	if err != nil {
		panic("Error getting pokemon")
	}
	return mapper.MapPokemon(pokemonResponse)
}