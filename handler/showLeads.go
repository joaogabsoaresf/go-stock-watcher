package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabsoaresf/mygocrm/schemas"
)

// @BasePath /api/v1

// @Summary Show leads
// @Description Show all leads in database
// @Tags Leads
// @Accept json
// @Produce json
// @Success 200 {object} ShowLeadResponse
// @Failures 500 {object} ErrorResponse
// @Router /leads [get]
func ShowLeadsHandler(ctx *gin.Context) {
	lead := []schemas.Lead{}

	if err := db.Find(&lead).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error find lead")
		return
	}

	sendSuccess(ctx, "list", lead)
}
