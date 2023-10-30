package router

import (
	"github.com/GuilhermeVozniak/go-opportunities/handlers"
	"github.com/gin-gonic/gin"
)

func routes(router *gin.Engine) {

	// Initialize handlers
	handlers.InitializeHandler()

	// maintenance
	router.GET("/health", handlers.GetHealthHandler)

	// api
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opportunities", handlers.GetAllOpportunitiesHandler)
		v1.GET("/opportunity/:id", handlers.GetOpportunityByIdHandler)
		v1.POST("/opportunity", handlers.CreateOpportunityHandler)
		v1.PUT("/opportunity/:id", handlers.UpdateOpportunityByIdHandler)
		v1.DELETE("/opportunity/:id", handlers.DeleteOpportunityByIdHandler)
	}

}
