package handlers

import (
	"errors"
	"fmt"
	opportunityRequestDTO "github.com/GuilhermeVozniak/go-opportunities/dto/opportunity/requests"
	opportunityResponseDTO "github.com/GuilhermeVozniak/go-opportunities/dto/opportunity/responses"
	validationHelpers "github.com/GuilhermeVozniak/go-opportunities/validationHelpers"
	"github.com/GuilhermeVozniak/go-opportunities/schemas"
	"github.com/GuilhermeVozniak/go-opportunities/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

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
	response = response.FromModel(&opportunity)

	utils.SendSuccessResponseWithCode(ctx, http.StatusOK, "update-opportunity", response)
}
