package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func routes(router *gin.Engine) {
	// maintenance
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"online": true,
		})
	})

	// api
	v1 := router.Group("/api/v1")
	{

		// get all opportunities
		v1.GET("/opportunities", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hello World",
			})
		})

		// get opportunity by id
		v1.GET("/opportunity/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hello World",
			})
		})

		// create all opportunity
		v1.POST("/opportunity", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hello World",
			})
		})

		// update opportunity by id
		v1.PUT("/opportunity/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hello World",
			})
		})

		// delete opportunity by id
		v1.DELETE("/opportunity/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hello World",
			})
		})
		

	}

}
