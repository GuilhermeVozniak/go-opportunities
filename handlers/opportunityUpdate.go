package handlers

import (
	"errors"
	"fmt"
	opportunityRequestDTO "github.com/GuilhermeVozniak/go-opportunities/dto/opportunity/requests"
	opportunityResponseDTO "github.com/GuilhermeVozniak/go-opportunities/dto/opportunity/responses"
	"github.com/GuilhermeVozniak/go-opportunities/schemas"
	"github.com/GuilhermeVozniak/go-opportunities/utils"
	validationHelpers "github.com/GuilhermeVozniak/go-opportunities/validationHelpers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// @BasePath /api/v1

// @Summary Update opportunity
// @Description Updates a job opportunity
// @Tags Opportunities
// @Accept  json
// @Produce  json
// @Param id path string true "Opportunity ID"
// @Param request body dto.UpdateOpportunityRequest true "Request body"
// @Success 200 {object} dto.UpdateOpportunityResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /opportunity [put]
func UpdateOpportunityByIdHandler(ctx *gin.Context) {
	var (
		id          = ctx.Param("id")
		request     opportunityRequestDTO.UpdateOpportunityRequest
		response    *opportunityResponseDTO.UpdateOpportunityResponse
		opportunity = schemas.Opportunity{}
		err         error
	)

	// check for id
	if id == "" {
		utils.SendErrorResponseWithCode(ctx, http.StatusBadRequest, validationHelpers.ErrParamIsRequired("id", "string"))
		return
	}

	// parse request
	if err = ctx.BindJSON(&request); err != nil {
		utils.SendErrorResponseWithCode(ctx, http.StatusBadRequest, err)
		return
	}
	// validate request
	if err = request.Validate(); err != nil {
		utils.SendErrorResponseWithCode(ctx, http.StatusBadRequest, err)
		return
	}

	// check for opportunity
	if err = db.First(&opportunity, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.SendErrorResponseWithCode(ctx, http.StatusNotFound, errors.New(fmt.Sprintf("opportunity with id %s not found", id)))
		} else {
			utils.SendErrorResponseWithCode(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	// convert dto to model
	opportunity = request.ToModel(&opportunity)

	if err := db.Updates(&opportunity).Error; err != nil {
		logger.Error(err)
		utils.SendErrorResponseWithCode(ctx, http.StatusInternalServerError, err)
		return
	}

	// convert model to dto
	response = response.FromModelToResponse(&opportunity, "update-opportunity")

	utils.SendSuccessResponseWithCode(ctx, http.StatusOK, &response)
}
