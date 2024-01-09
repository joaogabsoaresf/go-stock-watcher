package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabsoaresf/mygocrm/schemas"
)

// @BasePath /api/v1

// @Summary Show lead
// @Description Show a specifc lead in database
// @Tags Leads
// @Accept json
// @Produce json
// @Param id query string true "Lead id"
// @Success 200 {object} ShowLeadResponse
// @Failures 400 {object} ErrorResponse
// @Failures 404 {object} ErrorResponse
// @Router /lead [get]
func ShowLeadHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	lead := schemas.Lead{}
	// find lead
	if err := db.First(&lead, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("lead with id: %s not found", id))
		return
	}

	// return lead
	sendSuccess(ctx, "find", lead)
}
