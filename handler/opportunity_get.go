package handler

import (
	"errors"
	"fmt"
	"net/http"

	opportunityResponseDTO "github.com/GuilhermeVozniak/go-opportunities/dto/opportunity/responses"
	"github.com/GuilhermeVozniak/go-opportunities/helper"
	"github.com/GuilhermeVozniak/go-opportunities/schemas"
	"github.com/GuilhermeVozniak/go-opportunities/util"
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
// @Failure 400 {object} util.ErrorResponse
// @Failure 404 {object} util.ErrorResponse
// @Failure 500 {object} util.ErrorResponse
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
		util.SendErrorResponseWithCode(ctx, http.StatusBadRequest, helper.ErrParamIsRequired("id", "string"))
		return
	}

	// check for opportunity
	if err = db.First(&opportunity, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			util.SendErrorResponseWithCode(ctx, http.StatusNotFound, errors.New(fmt.Sprintf("opportunity with id %s not found", id)))
		} else {
			util.SendErrorResponseWithCode(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	// convert model to dto
	response = response.FromModelToResponse(&opportunity, "get-opportunity")

	util.SendSuccessResponseWithCode(ctx, http.StatusOK, &response)
}
