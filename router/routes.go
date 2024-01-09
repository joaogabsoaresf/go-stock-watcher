package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaogabsoaresf/go-stock-watcher/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/lead", handler.ShowLeadHandler)
		v1.POST("/lead", handler.CreateLeadHandler)
		v1.DELETE("/lead", handler.DeleteLeadHandler)
		v1.PUT("/lead", handler.UpdateLeadHandler)
		v1.GET("/leads", handler.ShowLeadsHandler)
	}
}
