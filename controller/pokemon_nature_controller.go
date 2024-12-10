package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mitchie-noc/gopoke-api/nature"
)

func GetNature(context *gin.Context) {
	name := context.Param("name")
	natureResponse := nature.GetNature(name)
	context.JSON(http.StatusOK, natureResponse)
}

func GetAllNatures(context *gin.Context) {
	natureResponse := nature.GetAllNatures()
	context.JSON(http.StatusOK, natureResponse)
}