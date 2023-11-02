package handler

import (
	opportunityRequestDTO "github.com/GuilhermeVozniak/go-opportunities/dto/opportunity/requests"
	opportunityResponseDTO "github.com/GuilhermeVozniak/go-opportunities/dto/opportunity/responses"
	"github.com/GuilhermeVozniak/go-opportunities/util"
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
// @Failure 400 {object} util.ErrorResponse
// @Failure 422 {object} util.ErrorResponse
// @Failure 500 {object} util.ErrorResponse
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
		util.SendErrorResponseWithCode(ctx, http.StatusBadRequest, err)
		return
	}
	// validate request
	if err = request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		util.SendErrorResponseWithCode(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	// convert dto to model
	opportunity := request.ToModel()

	if err := db.Create(&opportunity).Error; err != nil {
		logger.Error(err)
		util.SendErrorResponseWithCode(ctx, http.StatusInternalServerError, err)
		return
	}

	// convert model to dto
	response = response.FromModelToResponse(opportunity,"create-opportunity")
	util.SendSuccessResponseWithCode(ctx, http.StatusCreated, &response)
}
