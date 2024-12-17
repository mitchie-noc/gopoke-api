package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mitchie-noc/gopoke-api/calculator"
)

func GetPokemonTypeMatchup(context *gin.Context) {
	var request struct {
		Types []string `json:"types"`
	}

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if len(request.Types) < 1 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "At least one type is required"})
		return
	}

	matchUps, err := calculator.GetMatchups(request.Types)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, matchUps)
}