package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabsoaresf/go-stock-watcher/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("lead %s success", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message    string `json:"message"`
	ErrorCorde string `json:"errorCord"`
}

type CreateLeadResponse struct {
	Message string              `json:"message"`
	Data    schemas.LeadRespose `json:"data"`
}

type DeleteLeadResponse struct {
	Message string              `json:"message"`
	Data    schemas.LeadRespose `json:"data"`
}

type ShowLeadResponse struct {
	Message string              `json:"message"`
	Data    schemas.LeadRespose `json:"data"`
}
