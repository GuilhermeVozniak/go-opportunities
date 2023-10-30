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
