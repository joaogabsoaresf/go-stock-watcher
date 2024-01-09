package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabsoaresf/go-stock-watcher/schemas"
)

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
