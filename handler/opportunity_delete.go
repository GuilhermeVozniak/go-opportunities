package handler

import (
	"errors"
	"fmt"
	opportunityResponseDTO "github.com/GuilhermeVozniak/go-opportunities/dto/opportunity/responses"
	"github.com/GuilhermeVozniak/go-opportunities/schemas"
	"github.com/GuilhermeVozniak/go-opportunities/util"
	"github.com/GuilhermeVozniak/go-opportunities/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// @BasePath /api/v1

// @Summary Delete opportunity
// @Description Delete a job opportunity
// @Tags Opportunities
// @Accept  json
// @Produce  json
// @Param id path string true "Opportunity ID"
// @Success 200 {object} dto.DeleteOpportunityResponse
// @Failure 400 {object} util.ErrorResponse
// @Failure 404 {object} util.ErrorResponse
// @Failure 500 {object} util.ErrorResponse
// @Router /opportunity [delete]
func DeleteOpportunityByIdHandler(ctx *gin.Context) {
	var (
		id          string = ctx.Param("id")
		opportunity schemas.Opportunity
		response    *opportunityResponseDTO.DeleteOpportunityResponse
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

	// delete opportunity
	if err = db.Delete(&opportunity).Error; err != nil {
		util.SendErrorResponseWithCode(ctx, http.StatusInternalServerError, err)
		return
	}

	// convert model to dto
	response = response.FromModelToResponse(&opportunity, "delete-opportunity")
	util.SendSuccessResponseWithCode(ctx, http.StatusOK, &response)
}
