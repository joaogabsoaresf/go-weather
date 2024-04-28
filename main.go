package main

import (
	"ifoodchallenge/config"
	"ifoodchallenge/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Config init error: %v", err)
		return
	}
	router.Initialize()
}
