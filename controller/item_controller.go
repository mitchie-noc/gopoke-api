package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mitchie-noc/gopoke-api/item"
)

func GetItemsByCategory(context *gin.Context) {
	name := context.Query("category")
	itemResponse := item.GetItemsByCategory(name)
	context.JSON(http.StatusOK, itemResponse)
}

func GetItemByName(context *gin.Context) {
	name := context.Param("name")
	itemResponse := item.GetItemByName(name)
	context.JSON(http.StatusOK, itemResponse)
}