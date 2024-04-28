package handler

import (
	config "github.com/joaogabsoaresf/go-weather/config"
)

var (
	logger *config.Logger
	// secrets *config.Secrets
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	// secrets = config.GetSecrets()
}
