package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mitchie-noc/gopoke-api/pokemon"
)

func GetPokemon(context *gin.Context) {
	name := context.Param("name")
	pokemonRespone := pokemon.GetPokemon(name)
	context.JSON(http.StatusOK, pokemonRespone)
}