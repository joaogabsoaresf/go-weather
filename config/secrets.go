package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

const (
	ENV_FILE = "./dotenv-files/.env"
)

type Secrets struct {
	Env              string `env:"ENV"`
	OPEN_WEATHER_API string `env:"OPEN_WEATHER_API"`
}

func InitializeSecrets() *Secrets {
	logger = GetLogger("secrets")
	secrets := parseEnv()
	return secrets
}

func loadEnvFile() error {
	if _, err := os.Stat(ENV_FILE); os.IsNotExist(err) {
		return err
	}
	return godotenv.Load(ENV_FILE)
}

func parseEnv() *Secrets {
	if err := loadEnvFile(); err != nil {
		logger.Errorf("erro to load dotenv file, error: %v", err)
		return &Secrets{}
	}
	sct := &Secrets{}
	if err := env.Parse(sct); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return sct
}
