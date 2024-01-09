package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabsoaresf/go-stock-watcher/schemas"
)

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
