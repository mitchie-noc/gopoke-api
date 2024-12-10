package resource

import (
	"fmt"
	"strings"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

// Helper function to handle common logic for fetching resources
func GetResource(offset int, limit int, resourceName string) (structs.Resource){
	// Fetch resource data
	resourceResponse := getResourceResponse(resourceName, offset, limit)

	// Update Next URL
	resourceResponse.Next = strings.Replace(resourceResponse.Next, "https://pokeapi.co/api/v2", "http://localhost:8080/api/v1", 1)

	// Calculate Previous URL with proper offset
	previousOffset := offset - limit
	if previousOffset < 0 {
		previousOffset = 0
	}
	resourceResponse.Previous = fmt.Sprintf("http://localhost:8080/api/v1/%s?offset=%d&limit=%d", resourceName, previousOffset, limit)

	return resourceResponse
}

func SearchResource(term string, resource string) (structs.Resource){
	resourceResponse := searchResource(resource, term)
	return resourceResponse
}

func getResourceResponse(resource string, offset int, limit int) (structs.Resource) {
	result, err := pokeapi.Resource(resource, offset, limit)
	if err != nil {
		panic("Error getting resource")
	}
	return result
}

func searchResource(resource string, term string) (structs.Resource) {
	result, err := pokeapi.Search(resource, term)
	if err != nil {
		panic("Error searching for resource") 
	}
	return result
}

