package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/matheushermes/FinGO/configs"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var Rdb *redis.Client

func InitRedis() {
	db, _ := strconv.Atoi(configs.REDIS_DB)
	Rdb = redis.NewClient(&redis.Options{
		Addr: 		fmt.Sprintf("%s:%s", configs.REDIS_HOST, configs.REDIS_PORT),
		DB: 		db,
		Password: 	configs.REDIS_PASSWORD,
	})

	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	fmt.Println("Connected to Redis successfully")
}

func Set(key string, value interface{}, ttl time.Duration) error {
	return Rdb.Set(ctx, key, value, ttl).Err()
}

func Get(key string) (string, error) {
	return Rdb.Get(ctx, key).Result()
}