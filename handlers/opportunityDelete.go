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

// @Summary Delete opportunity
// @Description Delete a job opportunity
// @Tags Opportunities
// @Accept  json
// @Produce  json
// @Param id path string true "Opportunity ID"
// @Success 200 {object} dto.DeleteOpportunityResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
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

	// delete opportunity
	if err = db.Delete(&opportunity).Error; err != nil {
		utils.SendErrorResponseWithCode(ctx, http.StatusInternalServerError, err)
		return
	}

	// convert model to dto
	response = response.FromModelToResponse(&opportunity, "delete-opportunity")
	utils.SendSuccessResponseWithCode(ctx, http.StatusOK, &response)
}
