package router

import (
	helathHandlers "github.com/GuilhermeVozniak/go-opportunities/handlers/health"
	opportunitiesHandlers "github.com/GuilhermeVozniak/go-opportunities/handlers/opportunity"
	"github.com/gin-gonic/gin"
)

func routes(router *gin.Engine) {
	// maintenance
	router.GET("/health", helathHandlers.GetHealthHandler)

	// api
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opportunities", opportunitiesHandlers.GetAllOpportunitiesHandler)
		v1.GET("/opportunity/:id", opportunitiesHandlers.GetOpportunityByIdHandler)
		v1.POST("/opportunity", opportunitiesHandlers.CreateOpportunityHandler)
		v1.PUT("/opportunity/:id", opportunitiesHandlers.UpdateOpportunityByIdHandler)
		v1.DELETE("/opportunity/:id", opportunitiesHandlers.DeleteOpportunityByIdHandler)
	}

}
