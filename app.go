package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"ryantest.com/learningGo/service"
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

	server.Run("localhost:8080")
}

func getPokemonResource(context *gin.Context) {
	offset, limit, err := getOffsetAndLimit(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, ApiError{"GoPoke-0001", "Invalid query params for offset / limit"})
		return
	}

	resourceResponse := service.GetResource("pokemon", offset, limit)
	resourceResponse.Next = strings.Replace(resourceResponse.Next, "https://pokeapi.co/api/v2", "http://localhost:8080/api/v1",1)
	previousOffset := offset-limit
	if previousOffset < 0 {
		previousOffset = 0
	}
	resourceResponse.Previous = fmt.Sprintf("http://localhost:8080/api/v1/pokemon?offset=%d&limit=%d",previousOffset, limit)
	context.JSON(http.StatusOK, resourceResponse)
}

func getPokemon(context *gin.Context) {
	name := context.Param("name")
	pokemon := service.GetPokemon(name)
	context.JSON(http.StatusOK, pokemon)
	fmt.Println(pokemon)
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

