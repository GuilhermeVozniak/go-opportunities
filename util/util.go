package util

import (
	"github.com/gin-gonic/gin"
)

func SendErrorResponseWithCode(ctx *gin.Context, code int, err error) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(code, ErrorResponse{
		Message:   err.Error(),
		ErrorCode: code,
	})
}

func SendSuccessResponseWithCode(ctx *gin.Context, code int, data interface{}) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(code, data)
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

