package router

import (
	docs "github.com/GuilhermeVozniak/go-opportunities/docs"
	"github.com/GuilhermeVozniak/go-opportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func routes(router *gin.Engine) {
	var (
		basePath = "/api/v1"
	)

	// Initialize handlers
	handler.InitializeHandler()

	// Initialize swagger
	docs.SwaggerInfo.BasePath = basePath
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// maintenance
	router.GET("/health", handler.GetHealthHandler)

	// api
	v1 := router.Group(basePath)
	{
		v1.GET("/opportunities", handler.GetAllOpportunitiesHandler)
		v1.GET("/opportunity/:id", handler.GetOpportunityByIdHandler)
		v1.POST("/opportunity", handler.CreateOpportunityHandler)
		v1.PUT("/opportunity/:id", handler.UpdateOpportunityByIdHandler)
		v1.DELETE("/opportunity/:id", handler.DeleteOpportunityByIdHandler)
	}

}
