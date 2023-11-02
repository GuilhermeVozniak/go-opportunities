package handler

import (
	opportunityResponseDTO "github.com/GuilhermeVozniak/go-opportunities/dto/opportunity/responses"
	"github.com/GuilhermeVozniak/go-opportunities/schemas"
	"github.com/GuilhermeVozniak/go-opportunities/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary List all opportunity
// @Description List all job opportunity
// @Tags Opportunities
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.GetallOpportunityResponse
// @Failure 400 {object} util.ErrorResponse
// @Failure 404 {object} util.ErrorResponse
// @Failure 500 {object} util.ErrorResponse
// @Router /opportunities [get]
func GetAllOpportunitiesHandler(ctx *gin.Context) {
	var (
		opportunities = []schemas.Opportunity{}
		response      *opportunityResponseDTO.GetallOpportunityResponse
		err           error
	)

	if err = db.Find(&opportunities).Error; err != nil {
		logger.Error(err)
		util.SendErrorResponseWithCode(ctx, http.StatusInternalServerError, err)
		return
	}

	response = response.FromModelToResponse(&opportunities, "list-opportunities")
	util.SendSuccessResponseWithCode(ctx, http.StatusOK, &response)
}
