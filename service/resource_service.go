package service

import (
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetResource(resource string, offset int, limit int) (structs.Resource) {
	result, err := pokeapi.Resource(resource, offset, limit)
	if err != nil {
		panic("Error getting resource")
	}
	return result
}

func SearchResource(resource string, term string) (structs.Resource) {
	result, err := pokeapi.Search(resource, term)
	if err != nil {
		panic("Error searching for resource") 
	}
	return result
}