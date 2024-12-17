package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/mitchie-noc/gopoke-api/controller"
	"github.com/mitchie-noc/gopoke-api/calculator"
)

func main() {
	calculator.LoadTypeEffectiveness()

	server := gin.Default()
	server.Use(cors.Default())

	server.GET("/api/v1/pokemon", controller.GetPokemonResource)
	server.GET("/api/v1/pokemon/:name", controller.GetPokemon)
	server.GET("/api/v1/pokemon/search", controller.SearchPokemonResource)
	server.GET("/api/v1/nature", controller.GetAllNatures)
	server.GET("/api/v1/nature/:name", controller.GetNature)
	server.GET("/api/v1/item", controller.GetItemsByCategory)
	server.GET("/api/v1/item/:name", controller.GetItemByName)
	server.POST("/api/v1/type/matchup", controller.GetPokemonTypeMatchup)

	server.Run("localhost:8080")


}

