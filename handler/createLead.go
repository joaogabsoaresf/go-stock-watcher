package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabsoaresf/go-stock-watcher/schemas"
)

// @BasePath /api/v1

// @Summary Create lead
// @Description Create a new lead in database
// @Tags Leads
// @Accept json
// @Produce json
// @Param request body CreateLeadRequest true "Request body"
// @Success 200 {object} CreateLeadResponse
// @Failures 400 {object} ErrorResponse
// @Failures 500 {object} ErrorResponse
// @Router /lead [post]
func CreateLeadHandler(ctx *gin.Context) {
	request := CreateLeadRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	lead := schemas.Lead{
		Name:   request.Name,
		Email:  request.Email,
		Phone:  request.Phone,
		Client: *request.Client,
		Age:    request.Age,
	}

	if err := db.Create(&lead).Error; err != nil {
		logger.Errorf("lead creation error: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create", lead)
}
