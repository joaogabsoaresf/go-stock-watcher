package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowLeadHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Lead",
	})
}

func CreateLeadHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST Lead",
	})
}

func DeleteLeadHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Lead",
	})
}

func UpdateLeadHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Lead",
	})
}

func ShowLeadsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Leads",
	})
}
