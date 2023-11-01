package handlers

import (
	"errors"
	"fmt"
	"net/http"

	opportunityResponseDTO "github.com/GuilhermeVozniak/go-opportunities/dto/opportunity/responses"
	validationHelpers "github.com/GuilhermeVozniak/go-opportunities/validationHelpers"
	"github.com/GuilhermeVozniak/go-opportunities/schemas"
	"github.com/GuilhermeVozniak/go-opportunities/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @BasePath /api/v1

// @Summary Get job opportunity
// @Description Get a job opportunity details
// @Tags Opportunities
// @Accept  json
// @Produce  json
// @Param id path string true "Opportunity ID"
// @Success 200 {object} dto.GetOpportunityResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /opportunity [get]
func GetOpportunityByIdHandler(ctx *gin.Context) {
	var (
		id          string = ctx.Param("id")
		opportunity schemas.Opportunity
		response    *opportunityResponseDTO.GetOpportunityResponse
		err         error
	)

	// check for id
	if id == "" {
		utils.SendErrorResponseWithCode(ctx, http.StatusBadRequest, validationHelpers.ErrParamIsRequired("id", "string"))
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

	// convert model to dto
	response = response.FromModelToResponse(&opportunity, "get-opportunity")

	utils.SendSuccessResponseWithCode(ctx, http.StatusOK, &response)
}
