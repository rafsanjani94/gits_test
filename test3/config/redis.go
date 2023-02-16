package config

import (
	"context"
	"fmt"
	"gits/test3/conv"
	"os"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func ConnectRedis() {
	REDIS_HOST := os.Getenv("REDISV6_HOST")
	REDIS_PORT := conv.StringToInt(os.Getenv("REDISV6_PORT"), 0)
	REDIS_AUTH := os.Getenv("REDISV6_AUTH")
	REDIS_DB := conv.StringToInt(os.Getenv("REDISV6_DB"), 0)

	opt := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", REDIS_HOST, REDIS_PORT),
		Password: REDIS_AUTH,
		DB:       REDIS_DB,
	}

	RedisClient = redis.NewClient(opt)

	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
