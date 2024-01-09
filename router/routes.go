package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/joaogabsoaresf/mygocrm/docs"
	"github.com/joaogabsoaresf/mygocrm/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// init handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/lead", handler.ShowLeadHandler)
		v1.POST("/lead", handler.CreateLeadHandler)
		v1.DELETE("/lead", handler.DeleteLeadHandler)
		v1.PUT("/lead", handler.UpdateLeadHandler)
		v1.GET("/leads", handler.ShowLeadsHandler)
	}

	// init swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
