package redis_db

import (
	"context"
	"log"
	"todoflow-api/internal/config"

	"github.com/redis/go-redis/v9"
)

var (
	RDB *redis.Client
	Ctx = context.Background()
)

func Init(conf *config.Config) *redis.Client {
	if conf.Redis.Status == 1 {
		return RDB
	}
	RDB = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Address,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})

	_, err := RDB.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect Redis: %v", err)
		return RDB
	}
	log.Println("Connected to redis")
	return RDB
}
