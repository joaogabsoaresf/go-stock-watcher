package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabsoaresf/go-stock-watcher/schemas"
)

// @BasePath /api/v1

// @Summary Delete lead
// @Description Delete a lead in database
// @Tags Leads
// @Accept json
// @Produce json
// @Param request body CreateLeadRequest true "Request body"
// @Success 200 {object} DeleteLeadResponse
// @Failures 400 {object} ErrorResponse
// @Failures 404 {object} ErrorResponse
// @Failures 500 {object} ErrorResponse
// @Router /lead [delete]
func DeleteLeadHandler(ctx *gin.Context) {
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
	// delete lead
	if err := db.Delete(&lead).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting lead id: %s", id))
		return
	}
	sendSuccess(ctx, "delete", lead)
}
