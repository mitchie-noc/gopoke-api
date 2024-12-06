package service

import (
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetResource(resource string, offset int, limit int) (structs.Resource) {
	res, err := pokeapi.Resource(resource, offset, limit)
	if err != nil {
		panic("Error getting resource")
	}
	return res
}