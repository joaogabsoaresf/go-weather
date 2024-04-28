package handler

import (
	"fmt"
	"ifoodchallenge/connectors"

	"github.com/gin-gonic/gin"
)

func ReturnWeather(ctx *gin.Context) {
	city := ctx.Param("city")
	logger.Info(city)
	temp := getWeather(city)
	strTemp := fmt.Sprintf("%.1fºC", temp)
	sendSuccess(ctx, strTemp)
}

func getWeather(city string) float64 {
	api, err := connectors.GetWethear(city)
	if err != nil {
		logger.Errorf("Config init error: %v", err)
	}
	mainData := api["main"].(map[string]interface{})
	temp := mainData["temp"].(float64)
	return temp
}
