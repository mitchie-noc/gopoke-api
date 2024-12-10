package main

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/mitchie-noc/gopoke-api/resource"
	"github.com/mitchie-noc/gopoke-api/nature"
	"github.com/mitchie-noc/gopoke-api/pokemon"
)

type ApiError struct {
	Code string
	Message string
}

func main() {
	server := gin.Default()
	server.Use(cors.Default())

	server.GET("/api/v1/pokemon", getPokemonResource)
	server.GET("/api/v1/pokemon/:name", getPokemon)
	server.GET("/api/v1/pokemon/search", searchPokemonResource)
	server.GET("/api/v1/nature", getAllNatures)
	server.GET("/api/v1/nature/:name", getNature)

	server.Run("localhost:8080")
}

func getPokemon(context *gin.Context) {
	name := context.Param("name")
	pokemonRespone := pokemon.GetPokemon(name)
	context.JSON(http.StatusOK, pokemonRespone)
}

func getNature(context *gin.Context) {
	name := context.Param("name")
	natureResponse := nature.GetNature(name)
	context.JSON(http.StatusOK, natureResponse)
}

func getAllNatures(context *gin.Context) {
	natureResponse := nature.GetAllNatures()
	context.JSON(http.StatusOK, natureResponse)
}

func getPokemonNatures(context *gin.Context) {
	offset, limit, err := getOffsetAndLimit(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, ApiError{"GoPoke-0001", "Invalid query params for offset / limit"})
		return
	}
	resourceResponse := resource.GetResource(offset, limit, "nature")
	context.JSON(http.StatusOK, resourceResponse)
}

func getPokemonResource(context *gin.Context) {
	offset, limit, err := getOffsetAndLimit(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, ApiError{"GoPoke-0001", "Invalid query params for offset / limit"})
		return
	}
	resourceResponse := resource.GetResource(offset, limit, "pokemon")
	context.JSON(http.StatusOK, resourceResponse)
}

func searchPokemonResource(context *gin.Context) {
	searchTerm := context.Query("term")
	resourceResponse := resource.SearchResource(searchTerm, "pokemon")
	context.JSON(http.StatusOK, resourceResponse)
}

func getOffsetAndLimit(context *gin.Context) (int, int, error){
	offset := 0
	limit := 20

	providedOffset := context.Query("offset")
	providedLimit := context.Query("limit")

	if providedOffset != "" {
		fmt.Println("Updating offset to request value")
		newOffset, err := strconv.Atoi(providedOffset)

		if err != nil {
			return -1, -1, err
		}
		offset = newOffset
	}

	if providedLimit != "" {
		fmt.Println("Updating limit to request value")
		newLimit, err := strconv.Atoi(providedLimit)

		if err != nil {
			return -1, -1, err
		}
		limit = newLimit
	}
	return offset, limit, nil
}