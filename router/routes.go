package router

import (
	docs "github.com/GuilhermeVozniak/go-opportunities/docs"
	"github.com/GuilhermeVozniak/go-opportunities/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func routes(router *gin.Engine) {
	var (
		basePath = "/api/v1"
	)

	// Initialize handlers
	handlers.InitializeHandler()

	// Initialize swagger
	docs.SwaggerInfo.BasePath = basePath
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// maintenance
	router.GET("/health", handlers.GetHealthHandler)

	// api
	v1 := router.Group(basePath)
	{
		v1.GET("/opportunities", handlers.GetAllOpportunitiesHandler)
		v1.GET("/opportunity/:id", handlers.GetOpportunityByIdHandler)
		v1.POST("/opportunity", handlers.CreateOpportunityHandler)
		v1.PUT("/opportunity/:id", handlers.UpdateOpportunityByIdHandler)
		v1.DELETE("/opportunity/:id", handlers.DeleteOpportunityByIdHandler)
	}

}
