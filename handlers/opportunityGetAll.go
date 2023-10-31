package handlers

import (
	opportunityResponseDTO "github.com/GuilhermeVozniak/go-opportunities/dto/opportunity/responses"
	"github.com/GuilhermeVozniak/go-opportunities/schemas"
	"github.com/GuilhermeVozniak/go-opportunities/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllOpportunitiesHandler(ctx *gin.Context) {
	var (
		opportunities = []schemas.Opportunity{}
		response      *[]opportunityResponseDTO.GetallOpportunityResponse
		dto           opportunityResponseDTO.GetallOpportunityResponse
		err           error
	)

	if err = db.Find(&opportunities).Error; err != nil {
		logger.Error(err)
		utils.SendErrorResponseWithCode(ctx, http.StatusInternalServerError, err)
		return
	}

	response = dto.FromModel(&opportunities)

	utils.SendSuccessResponseWithCode(ctx, http.StatusOK, "list-opportunities", response)
}
