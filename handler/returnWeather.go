package handler

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	config "github.com/joaogabsoaresf/go-weather/config"
	connectors "github.com/joaogabsoaresf/go-weather/connectors"
)

func ReturnWeather(ctx *gin.Context) {
	city := ctx.Param("city")
	c := cityStringLower(city)
	cachedTemp := getCachedWeather(c)
	logger.Debug(cachedTemp)
	if cachedTemp != "" {
		sendSuccess(ctx, cachedTemp)
		return
	}
	temp := getWeather(c)
	strTemp := fmt.Sprintf("%.1fºC", temp)
	setCachedWeather(c, strTemp)
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

func cityStringLower(s string) string {
	return strings.ToLower(s)
}

func getCachedWeather(c string) string {
	v, err := config.GetData(*rdb, c)
	if err == fmt.Errorf("chave não encontrada") {
		return ""
	} else if err != nil {
		return ""
	}
	return v
}

func setCachedWeather(c string, temp string) {
	config.SetData(*rdb, c, temp)
}
