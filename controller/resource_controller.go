package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mitchie-noc/gopoke-api/resource"
	"github.com/mitchie-noc/gopoke-api/utils"
)

func GetPokemonResource(context *gin.Context) {
	offset, limit, err := utils.GetOffsetAndLimit(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.ApiError{Code: "GoPoke-0001", Message: "Invalid query params for offset / limit"})
		return
	}
	resourceResponse := resource.GetResource(offset, limit, "pokemon")
	context.JSON(http.StatusOK, resourceResponse)
}

func SearchPokemonResource(context *gin.Context) {
	searchTerm := context.Query("term")
	resourceResponse := resource.SearchResource(searchTerm, "pokemon")
	context.JSON(http.StatusOK, resourceResponse)
}

