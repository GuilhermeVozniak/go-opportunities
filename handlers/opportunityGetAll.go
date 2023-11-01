package handlers

import (
	opportunityResponseDTO "github.com/GuilhermeVozniak/go-opportunities/dto/opportunity/responses"
	"github.com/GuilhermeVozniak/go-opportunities/schemas"
	"github.com/GuilhermeVozniak/go-opportunities/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary List all opportunity
// @Description List all job opportunity
// @Tags Opportunities
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.DeleteOpportunityResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /opportunity [delete]
func GetAllOpportunitiesHandler(ctx *gin.Context) {
	var (
		opportunities = []schemas.Opportunity{}
		response      *opportunityResponseDTO.GetallOpportunityResponse
		dto           opportunityResponseDTO.GetallOpportunityResponse
		err           error
	)

	if err = db.Find(&opportunities).Error; err != nil {
		logger.Error(err)
		utils.SendErrorResponseWithCode(ctx, http.StatusInternalServerError, err)
		return
	}

	response = dto.FromModelToResponse(&opportunities, "list-opportunities")
	utils.SendSuccessResponseWithCode(ctx, http.StatusOK, &response)
}
