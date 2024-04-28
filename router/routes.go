package router

import (
	handler "github.com/joaogabsoaresf/go-weather/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// init handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("/weather/:city", handler.ReturnWeather)
	}
}
