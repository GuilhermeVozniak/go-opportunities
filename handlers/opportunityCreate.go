package handlers

import (
	opportunityRequestDTO "github.com/GuilhermeVozniak/go-opportunities/dto/opportunity/requests"
	opportunityResponseDTO "github.com/GuilhermeVozniak/go-opportunities/dto/opportunity/responses"
	"github.com/GuilhermeVozniak/go-opportunities/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary Create opportunity
// @Description Create a new job opportunity
// @Tags Opportunities
// @Accept  json
// @Produce  json
// @Param request body dto.CreateOpportunityRequest true "Reques body"
// @Success 201 {object} dto.CreateOpportunityResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 422 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /opportunity [post]
func CreateOpportunityHandler(ctx *gin.Context) {
	var (
		request  opportunityRequestDTO.CreateOpportunityRequest
		response *opportunityResponseDTO.CreateOpportunityResponse
		err      error
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
	response = response.FromModelToResponse(opportunity,"create-opportunity")
	utils.SendSuccessResponseWithCode(ctx, http.StatusCreated, &response)
}
