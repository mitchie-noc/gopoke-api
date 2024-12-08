package nature

import (
	"github.com/mtslzr/pokeapi-go"
	"github.com/mitchie-noc/gopoke-api/resource"
)

func GetAllNatures() ([]Nature){
	natureResources := resource.GetResource(0,100,"nature")
	var mappedNatures []Nature

	for _, natureResource := range natureResources.Results {
		mappedNature := getMappedNature(natureResource.Name)
		mappedNatures = append(mappedNatures, mappedNature)
	}

	return mappedNatures
}

func GetNature(name string) (Nature) {
	return getMappedNature(name)
}

func getMappedNature(nature string) (Nature) {	
	natureResponse, err := pokeapi.Nature(nature)
	if err != nil {
		panic("Error getting nature")
	}
	return MapNature(natureResponse)
}