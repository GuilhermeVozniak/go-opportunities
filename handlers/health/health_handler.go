package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetHealthHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"online": true,
	})
}