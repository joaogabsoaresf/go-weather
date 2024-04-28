package handler

import (
	config "github.com/joaogabsoaresf/go-weather/config"
	"github.com/redis/go-redis/v9"
)

var (
	logger *config.Logger
	rdb    *redis.Client
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	rdb = config.InitRedis()
}
