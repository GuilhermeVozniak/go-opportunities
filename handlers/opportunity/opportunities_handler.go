package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOpportunitiesHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func GetOpportunityByIdHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func CreateOpportunityHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func UpdateOpportunityByIdHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func DeleteOpportunityByIdHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}


