package cache

import (
	"Tigang/conf"
	"fmt"

	"github.com/go-redis/redis"
)

var (
	RedisClient *redis.Client
)

func InitRedis(){
	rcfg := conf.InitRedis()

	client := redis.NewClient(&redis.Options{
		Addr: rcfg.RedisAddr,
		Password: rcfg.RedisPw,
		DB: rcfg.RedisDbName,
	})

	_, err := client.Ping().Result()
	if err != nil{
		panic(fmt.Errorf("redis connect failed: %w", err))
	}
	RedisClient = client
}