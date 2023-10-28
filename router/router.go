package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeServer(){
	// Initialize the routes using gin default settings
	router := gin.Default()

	// Initialize the routes
	routes(router)

	// listen and serve on 0.0.0.0:8080
	router.Run("0.0.0.0:8080")
}