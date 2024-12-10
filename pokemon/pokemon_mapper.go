package pokemon

import (
	"github.com/mitchie-noc/gopoke-api/ability"
	"github.com/mtslzr/pokeapi-go/structs"
)

type Pokemon struct {
	Name string
	Sprite string
	Stats []Stat
	Types []Type
	Abilities []ability.Ability
}

type Stat struct {
	Name string
	Base int
	Effort int
	isBattleStat bool
}

type Type struct {
	Slot int
	Name string
}

func MapPokemon(original structs.Pokemon) Pokemon{
	pokemon := Pokemon{}
	pokemon.Name = original.Name
	pokemon.Sprite = original.Sprites.FrontDefault
	mapPokemonStats(original, &pokemon)
	mapPokemonTypes(original, &pokemon)
	mapPokemonAbilities(original, &pokemon)
	return pokemon
}

func mapPokemonAbilities(original structs.Pokemon, target *Pokemon) {
	for _, element := range original.Abilities {
		ability := ability.GetAbility(element.Ability.Name)
		target.Abilities = append(target.Abilities, ability)
	}
}

func mapPokemonStats(original structs.Pokemon, target *Pokemon) {
	for _, element := range original.Stats {
		stat := Stat {
			element.Stat.Name, 
			element.BaseStat, 
			element.Effort, 
			isBattleStat(element.Stat.Name)}
		target.Stats = append(target.Stats, stat)
	}
}

func mapPokemonTypes(original structs.Pokemon, target *Pokemon) {
	for _, element := range original.Types {
		pokemonType := Type {
			element.Slot, 
			element.Type.Name}
		target.Types = append(target.Types, pokemonType)
	}
}

func isBattleStat(name string) bool {
	return name != "hp"
}