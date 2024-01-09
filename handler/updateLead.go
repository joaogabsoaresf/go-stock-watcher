package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateLeadHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Lead",
	})
}
