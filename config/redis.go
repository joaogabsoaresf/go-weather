package config

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client
)

func InitRedis() *redis.Client {
	logger = GetLogger("redis")
	opt, err := redis.ParseURL("rediss://red-conb7ba1hbls73fg19eg:TkClebGBYvffeMTnNMCQiqLQ4HLBlmor@oregon-redis.render.com:6379")
	if err != nil {
		panic(err)
	}
	rdb = redis.NewClient(opt)
	logger.Info("Redis Connection Success")
	return rdb
}

func SetData(client redis.Client, k string, v string) {
	ctx := context.Background()
	err := client.Set(ctx, k, v, 10*time.Minute).Err()
	if err != nil {
		panic(err)
	}
}

func GetData(client redis.Client, k string) (string, error) {
	ctx := context.Background()
	val, err := client.Get(ctx, k).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("chave não encontrada")
	} else if err != nil {
		return "", err
	}
	return val, nil
}
