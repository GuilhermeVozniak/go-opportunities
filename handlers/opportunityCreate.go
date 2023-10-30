package handlers

import (
	"net/http"

	opportunityRequestDTO "github.com/GuilhermeVozniak/go-opportunities/dto/opportunity/requests"
	opportunityResponseDTO "github.com/GuilhermeVozniak/go-opportunities/dto/opportunity/responses"
	"github.com/GuilhermeVozniak/go-opportunities/utils"
	"github.com/gin-gonic/gin"
)

func CreateOpportunityHandler(ctx *gin.Context) {
	var (
		request opportunityRequestDTO.CreateOpportunityRequest
		response opportunityResponseDTO.OpportunityResponse
		err     error
	)

	// parse request
	if err = ctx.BindJSON(&request); err != nil {
		logger.Errorf("bind json error: %v", err)
		utils.SendErrorResponseWithCode(ctx, http.StatusBadRequest, err)
		return
	}
	// validate request
	if err = request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		utils.SendErrorResponseWithCode(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	// convert dto to model
	opportunity := request.ToModel()

	if err := db.Create(&opportunity).Error; err != nil {
		logger.Error(err)
		utils.SendErrorResponseWithCode(ctx, http.StatusInternalServerError, err)
		return
	}

	// convert model to dto
	response.FromModel(opportunity)

	utils.SendSuccessResponseWithCode(ctx, http.StatusCreated, "create-opportunity", opportunity)
}
