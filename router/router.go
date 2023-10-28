package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeServer(){
	// Initialize the routes using gin default settings
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"online": true,
		})
	})

	// listen and serve on 0.0.0.0:8080
	router.Run("0.0.0.0:8080")
}