package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func UpdateOpportunityByIdHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
