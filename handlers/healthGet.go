package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary Get application health
// @Description Return api heath status
// @Produce  json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func GetHealthHandler(ctx *gin.Context) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(http.StatusOK, HealthResponse{
		Message: "API is online",
		Online:  true,
	})
}

type HealthResponse struct {
	Message string `json:"message"`
	Online  bool   `json:"online"`
}
