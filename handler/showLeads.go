package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabsoaresf/go-stock-watcher/schemas"
)

func ShowLeadsHandler(ctx *gin.Context) {
	lead := []schemas.Lead{}

	if err := db.Find(&lead).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error find lead")
		return
	}

	sendSuccess(ctx, "list", lead)
}
