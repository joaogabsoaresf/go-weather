package handler

import (
	config "ifoodchallenge/config"
)

var (
	logger *config.Logger
	// secrets *config.Secrets
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	// secrets = config.GetSecrets()
}
