package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SendErrorResponseWithCode(ctx *gin.Context, code int, err error) {
	ctx.Header("content-type", "application/json")
	ctx.AbortWithStatusJSON(code, gin.H{
		"message":   err.Error(),
		"errorCode": code,
	})
}

func SendSuccessResponseWithCode(ctx *gin.Context, code int,op string, data interface{}) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": fmt.Sprintf("operation from handler %s successfull", op),
		"data":    data,
	})
}
